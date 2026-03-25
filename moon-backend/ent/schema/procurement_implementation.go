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

// ProcurementImplementation holds the schema definition for the ProcurementImplementation entity.
type ProcurementImplementation struct {
	ent.Schema
}

// Fields of the ProcurementImplementation.
func (ProcurementImplementation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_implementation_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementImplementationName").
			StorageKey("procurement_implementation_name").Annotations(entgql.OrderField("PROCUREMENT_IMPLEMENTATION_NAME")),

		field.String("ProcurementImplementationCode").
			StorageKey("procurement_implementation_code").Annotations(entgql.OrderField("PROCUREMENT_IMPLEMENTATION_CODE")),

		field.String("ProcurementImplementationDescription").
			StorageKey("procurement_implementation_description").Annotations(entgql.OrderField("PROCUREMENT_IMPLEMENTATION_DESCRIPTION")),

		field.String("ProcurementImplementationFlag").
			StorageKey("procurement_implementation_flag").Annotations(entgql.OrderField("PROCUREMENT_IMPLEMENTATION_FLAG")),

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

		field.UUID("procurement_plan_id", uuid.UUID{}).
			StorageKey("procurement_plan_id").Optional().Nillable(),

		field.UUID("organization_id", uuid.UUID{}).
			StorageKey("organization_id").Optional().Nillable(),

		field.UUID("user_id", uuid.UUID{}).
			StorageKey("user_id").Optional().Nillable(),
	}
}

// Edges of the ProcurementImplementation.
func (ProcurementImplementation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("procurement_plan", ProcurementPlan.Type).
			Ref("procurement_implementations").
			Field("procurement_plan_id").
			Unique(),

		// 这里必须和反向边名字一致
		edge.To("procurement_acceptance", ProcurementAcceptance.Type).Unique(),
		edge.To("procurement_review", ProcurementReview.Type).Unique(),
	}
}

// Annotations of the ProcurementImplementation.
func (ProcurementImplementation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_implementation"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
