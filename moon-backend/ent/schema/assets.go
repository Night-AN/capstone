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

type Assets struct {
	ent.Schema
}

func (Assets) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("asset_id").
			Annotations(
				entgql.OrderField("ID")),

		field.String("asset_name").
			StorageKey("asset_name").Annotations(
			entgql.OrderField("ASSET_NAME")),

		field.String("asset_code").
			StorageKey("asset_code").
			Unique().Annotations(
			entgql.OrderField("ASSET_CODE")),

		field.String("asset_description").
			StorageKey("asset_description").
			Default("").Annotations(
			entgql.OrderField("ASSET_DESCRIPTION")),

		field.String("asset_flag").
			StorageKey("asset_flag").Annotations(
			entgql.OrderField("ASSET_FLAG")),

		field.Int("quantity").
			StorageKey("quantity").
			Default(1).Annotations(
			entgql.OrderField("QUANTITY")),

		field.String("location").
			StorageKey("location").
			Default("").Annotations(
			entgql.OrderField("LOCATION")),

		field.Float("purchase_price").
			StorageKey("purchase_price").Annotations(
			entgql.OrderField("PURCHASE_PRICE")),

		field.Float("depreciation_price").
			StorageKey("depreciation_price").Annotations(
			entgql.OrderField("DEPRECIATION_PRICE")),

		field.Time("purchase_date").
			StorageKey("purchase_date").
			Optional().Nillable().Annotations(
			entgql.OrderField("PURCHASE_DATE")),

		field.String("manufacturer").
			StorageKey("manufacturer").
			Default("").Annotations(
			entgql.OrderField("MANUFACTURER")),

		field.String("model").
			StorageKey("model").
			Default("").Annotations(
			entgql.OrderField("MODEL")),

		field.JSON("other_metadata", map[string]interface{}{}).
			StorageKey("other_metadata").
			Optional(),

		field.UUID("asset_type_id", uuid.UUID{}).
			StorageKey("asset_type_id").
			Optional().Nillable().
			Annotations(entgql.OrderField("ASSET_TYPE_ID")),

		field.UUID("asset_category_id", uuid.UUID{}).
			StorageKey("asset_category_id").
			Optional().Nillable().
			Annotations(entgql.OrderField("ASSET_CATEGORY_ID")),
		field.UUID("organization_id", uuid.UUID{}).
			StorageKey("organization_id").
			Optional().Nillable().
			Annotations(entgql.OrderField("ORGANIZATION_ID")),

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

func (Assets) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("asset_category", AssetCategory.Type).
			Ref("asset").Field("asset_category_id").
			Unique(),
		edge.From("asset_type", AssetType.Type).
			Ref("asset").Field("asset_type_id").
			Unique(),
		edge.From("organization", Organization.Type).
			Ref("asset").Field("organization_id").
			Unique(),
	}
}

func (Assets) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "assets"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
