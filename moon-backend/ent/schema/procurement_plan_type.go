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

// ProcurementPlanType holds the schema definition for the ProcurementPlanType entity.
type ProcurementPlanType struct {
	ent.Schema
}

// Fields of the ProcurementPlanType.
func (ProcurementPlanType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_plan_type_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementPlanTypeName").
			StorageKey("procurement_plan_type_name").Annotations(entgql.OrderField("PROCUREMENT_PLAN_TYPE_NAME")),

		field.String("ProcurementPlanTypeCode").
			StorageKey("procurement_plan_type_code").Annotations(entgql.OrderField("PROCUREMENT_PLAN_TYPE_CODE")),

		field.String("ProcurementPlanTypeFlag").
			StorageKey("procurement_plan_type_flag").Annotations(entgql.OrderField("PROCUREMENT_PLAN_TYPE_FLAG")),

		field.Time("CreatedAt").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),

		field.Time("UpdatedAt").
			StorageKey("updated_at").
			Optional().
			Nillable().
			Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the ProcurementPlanType.
func (ProcurementPlanType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("procurement_plan", ProcurementPlan.Type),
	}
}

// Annotations of the ProcurementPlanType.
func (ProcurementPlanType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_plan_type"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
