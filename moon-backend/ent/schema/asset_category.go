package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type AssetCategory struct {
	ent.Schema
}

func (AssetCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("asset_category_id").
			Annotations(
				entgql.OrderField("ID")),

		field.String("category_name").
			StorageKey("category_name").
			Unique().Annotations(
			entgql.OrderField("CATEGORY_NAME")),

		field.String("category_code").
			StorageKey("category_code").
			Unique().Annotations(
			entgql.OrderField("CATEGORY_CODE")),

		field.String("category_flag").
			StorageKey("category_flag").
			Default("ACTIVE"),

		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),

		field.Time("updated_at").
			StorageKey("updated_at").
			Optional().Nillable().
			Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (AssetCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("asset_type", AssetType.Type),
		edge.To("asset", Assets.Type),
	}
}

func (AssetCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "asset_category"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
