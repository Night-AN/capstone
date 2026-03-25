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

// ProcurementPlan holds the schema definition for the ProcurementPlan entity.
type ProcurementPlan struct {
	ent.Schema
}

// Fields of the ProcurementPlan.
func (ProcurementPlan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_plan_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementPlanName").
			StorageKey("procurement_plan_name").Annotations(entgql.OrderField("PROCUREMENT_PLAN_NAME")),

		field.String("ProcurementPlanCode").
			StorageKey("procurement_plan_code").Annotations(entgql.OrderField("PROCUREMENT_PLAN_CODE")),

		field.String("ProcurementPlanDescription").
			StorageKey("procurement_plan_description").Annotations(entgql.OrderField("PROCUREMENT_PLAN_DESCRIPTION")),

		field.String("ProcurementPlanFlag").
			StorageKey("procurement_plan_flag").Annotations(entgql.OrderField("PROCUREMENT_PLAN_FLAG")),

		field.Int("ProcurementPlanQuantity").
			StorageKey("procurement_plan_quantity").
			Default(1).Annotations(entgql.OrderField("PROCUREMENT_PLAN_QUANTITY")),

		field.Float("ProcurementPlanPrice").
			StorageKey("procurement_plan_price").Annotations(entgql.OrderField("PROCUREMENT_PLAN_PRICE")),

		field.Time("ProcurementPlanPurchaseDate").
			StorageKey("procurement_plan_purchase_date").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("PROCUREMENT_PLAN_PURCHASE_DATE")),

		field.String("ProcurementPlanPurchaseType").
			StorageKey("procurement_plan_purchase_type").
			Default("").Annotations(entgql.OrderField("PROCUREMENT_PLAN_PURCHASE_TYPE")),

		field.JSON("OtherMetadata", map[string]interface{}{}).
			StorageKey("other_metadata").
			Default(map[string]interface{}{}),

		field.Time("CreatedAt").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),

		field.Time("UpdatedAt").
			StorageKey("updated_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("UPDATED_AT")),

		field.UUID("procurement_plan_type_id", uuid.UUID{}).
			StorageKey("procurement_plan_type_id").Optional().Nillable(),

		field.UUID("organization_id", uuid.UUID{}).
			StorageKey("organization_id").Optional().Nillable(),

		field.UUID("user_id", uuid.UUID{}).
			StorageKey("user_id").Optional().Nillable(),
	}
}

// Edges of the ProcurementPlan.
func (ProcurementPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("procurement_plan_type", ProcurementPlanType.Type).Ref("procurement_plan").
			Field("procurement_plan_type_id").
			Unique(),

		edge.To("procurement_implementations", ProcurementImplementation.Type),
	}
}

// Annotations of the ProcurementPlan.
func (ProcurementPlan) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_plan"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
