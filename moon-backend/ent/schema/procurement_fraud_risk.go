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

type ProcurementFraudRisk struct {
	ent.Schema
}

func (ProcurementFraudRisk) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("risk_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),

		field.UUID("procurement_plan_id", uuid.UUID{}).
			StorageKey("procurement_plan_id").Optional().Nillable(),

		field.UUID("procurement_implementation_id", uuid.UUID{}).
			StorageKey("procurement_implementation_id").Optional().Nillable(),

		field.String("risk_level").
			StorageKey("risk_level").Annotations(
			entgql.OrderField("RISK_LEVEL")),

		field.String("risk_reason").
			StorageKey("risk_reason").Optional().Nillable(),

		field.Int("risk_score").
			StorageKey("risk_score").Default(0),

		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

func (ProcurementFraudRisk) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (ProcurementFraudRisk) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("ai", "procurement_fraud_risk"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
