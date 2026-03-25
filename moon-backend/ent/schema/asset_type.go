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

type AssetType struct {
	ent.Schema
}

func (AssetType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("asset_type_id").
			Annotations(
				entgql.OrderField("ID")),

		field.String("asset_type_name").
			StorageKey("asset_type_name").Annotations(
			entgql.OrderField("ASSET_TYPE_NAME")),

		field.String("asset_type_code").
			StorageKey("asset_type_code").Annotations(
			entgql.OrderField("ASSET_TYPE_CODE")),

		field.String("asset_type_flag").
			StorageKey("asset_type_flag").
			Default("ACTIVE").Annotations(
			entgql.OrderField("ASSET_TYPE_FLAG")),

		field.UUID("asset_category_id", uuid.UUID{}).
			StorageKey("asset_category_id").
			Annotations(entgql.OrderField("ASSET_CATEGORY_ID")),

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

func (AssetType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset_category", AssetCategory.Type).
			Ref("asset_type").Field("asset_category_id").
			Unique().Required(),
		edge.To("asset", Assets.Type),
	}
}

func (AssetType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "asset_type"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
