package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserID         uuid.UUID
	OrganizationID uuid.UUID
	Permission     []string
	HasAny         bool
	jwt.RegisteredClaims
}

var jwtSecret = []byte("your-secret-key")

func GenerateToken(user_id, organization_id uuid.UUID, permission []string) (string, error) {
	claims := &UserClaims{
		UserID:         user_id,
		OrganizationID: organization_id,
		Permission:     permission,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyToken(tokenString string) (UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return UserClaims{}, errors.New("invalid token")
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return *claims, nil
	}
	return UserClaims{}, errors.New("invalid token")
}

func GetUserClaimsFromCtx(ctx context.Context) (*UserClaims, error) {
	// 直接获取你存在 ctx 里的 &claims
	claims, ok := ctx.Value("user_claims").(*UserClaims)
	if !ok || claims == nil {
		return nil, errors.New("用户未登录或token无效")
	}
	return claims, nil
}
