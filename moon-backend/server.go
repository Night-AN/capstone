package main

import (
	"context"
	"fmt"
	"log"
	"moon/ent"
	"moon/graph"
	"moon/pkg/auth"
	"moon/pkg/s3"
	"net/http"
	"os"
	"strings"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

const RUSTFS_ACCESSKEY = "7EubftlJSGnVZm4C8vKY"
const RUSTFS_SECRETKEY = "SwyYG39BnCPDsMTduvkZeilgEtfjQI20741HKraX"
const RUSTFS_API = "s3v4"
const RUSTFS_PATH = "auto"
const RUSTFS_RIGION = "cn-south-1"
const RUSTFS_ENDPOINT = "http://localhost:9000"
const RUSTFS_BUCKET = "capstone"

func main() {
	s3Client := s3.NewS3Client(context.Background(), s3.Config{
		Region:     RUSTFS_RIGION,
		Endpoint:   RUSTFS_ENDPOINT,
		AccessKey:  RUSTFS_ACCESSKEY,
		SecretKey:  RUSTFS_SECRETKEY,
		BucketName: RUSTFS_BUCKET,
	})

	client, err := ent.Open("postgres", "host=localhost port=5432 user=capstone password=capstone dbname=capstone sslmode=disable")
	if err != nil {
		log.Fatalf("opening connection: %v", err)
	}
	defer client.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	c := graph.Config{Resolvers: &graph.Resolver{Client: client, S3Client: s3Client}}
	c.Directives.HasPermission = func(ctx context.Context, obj any, next graphql.Resolver, required []string) (res any, err error) {
		// 从上下文中获取user_claims
		userClaims, ok := ctx.Value("user_claims").(*auth.UserClaims) // 假设UserClaims是你的JWT结构体类型
		if !ok || userClaims == nil {
			return nil, fmt.Errorf("未找到用户认证信息")
		}

		// 获取用户的权限列表
		userPermissions := userClaims.Permission
		if len(userPermissions) == 0 {
			return nil, fmt.Errorf("用户没有任何权限")
		}

		// 将用户权限转换为map以便快速查找
		permMap := make(map[string]bool)
		for _, perm := range userPermissions {
			if perm == "any::any" {
				return next(ctx)
			}
			permMap[perm] = true
		}

		// 逐一比对所需权限
		for _, reqPerm := range required {
			if !permMap[reqPerm] {
				return nil, fmt.Errorf("缺少必要权限: %s", reqPerm)
			}
		}
		return next(ctx)
	}
	srv := handler.New(graph.NewExecutableSchema(c))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	srv.Use(entgql.Transactioner{TxOpener: client})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", AuthMiddleWare()(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, CorsMiddleware(http.DefaultServeMux)))
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
	Data   interface{}   `json:"data"`
}

type ErrorDetail struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func AuthMiddleWare() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			// 默认：零值用户（匿名/未登录）
			claims := auth.UserClaims{
				UserID:         uuid.Nil,
				OrganizationID: uuid.Nil,
				Permission:     []string{},
			}

			// 1. 提取 Authorization 头
			token := r.Header.Get("Authorization")
			if token != "" {
				// 2. 校验 Bearer 格式
				parts := strings.Split(token, " ")
				if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
					tokenStr := parts[1]

					// 3. 验证 Token
					parsedClaims, err := auth.VerifyToken(tokenStr)
					if err == nil {
						// 只有验证成功，才覆盖 claims
						claims = parsedClaims
					}
				}
			}

			// 4. 无论如何都放入上下文，永远放行
			ctx := context.WithValue(r.Context(), "user_claims", &claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "false")

		next.ServeHTTP(w, r)
	})
}
