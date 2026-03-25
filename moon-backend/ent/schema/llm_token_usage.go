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

type LlmTokenUsage struct {
	ent.Schema
}

func (LlmTokenUsage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("usage_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),

		field.UUID("model_id", uuid.UUID{}).
			StorageKey("model_id"),

		field.UUID("user_id", uuid.UUID{}).
			StorageKey("user_id").Optional().Nillable(),

		field.Int("prompt_tokens").
			StorageKey("prompt_tokens").Default(0),

		field.Int("completion_tokens").
			StorageKey("completion_tokens").Default(0),

		field.Int("total_tokens").
			StorageKey("total_tokens").Default(0),

		field.Float("cost_amount").
			StorageKey("cost_amount").Default(0),

		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
	}
}

func (LlmTokenUsage) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (LlmTokenUsage) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("ai", "llm_token_usage"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
