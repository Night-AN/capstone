//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graph/ent.graphqls"),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithRelaySpec(true),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(func(g *gen.Graph, s *ast.Schema) error {
			uuidScalar := &ast.Definition{
				Name: "UUID",     // 标量名称
				Kind: ast.Scalar, // 类型为标量
			}
			if _, exists := s.Types["UUID"]; !exists {
				// 插入到第一个位置，确保其他类型能引用到
				s.Types["UUID"] = uuidScalar
			}
			return nil
		}),
		entgql.WithSchemaHook(func(g *gen.Graph, s *ast.Schema) error {
			timeScalar := &ast.Definition{
				Name: "Time",     // 标量名称
				Kind: ast.Scalar, // 类型为标量
			}
			if _, exists := s.Types["Time"]; !exists {
				// 插入到第一个位置，确保其他类型能引用到
				s.Types["Time"] = timeScalar
			}
			return nil
		}),
		entgql.WithSchemaHook(func(g *gen.Graph, s *ast.Schema) error {
			mapScalar := &ast.Definition{
				Name: "Map",      // 标量名称
				Kind: ast.Scalar, // 类型为标量
			}
			if _, exists := s.Types["Map"]; !exists {
				// 插入到第一个位置，确保其他类型能引用到
				s.Types["Map"] = mapScalar
			}
			return nil
		}),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{

		Features: []gen.Feature{
			gen.FeatureUpsert,
		},
		Templates: entgql.AllTemplates,
	}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
