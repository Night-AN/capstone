package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type LlmProcurementAnalysis struct {
	ent.Schema
}

func (LlmProcurementAnalysis) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("log_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),

		field.UUID("model_id", uuid.UUID{}).
			StorageKey("model_id").Optional().Nillable(),

		field.UUID("procurement_plan_id", uuid.UUID{}).
			StorageKey("procurement_plan_id").Optional().Nillable(),

		field.UUID("procurement_implementation_id", uuid.UUID{}).
			StorageKey("procurement_implementation_id").Optional().Nillable(),

		field.Text("analysis_result").
			StorageKey("analysis_result").Optional().Nillable(),

		field.UUID("token_usage_id", uuid.UUID{}).
			StorageKey("token_usage_id").Optional().Nillable(),

		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

func (LlmProcurementAnalysis) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (LlmProcurementAnalysis) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("ai", "llm_procurement_analysis"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
