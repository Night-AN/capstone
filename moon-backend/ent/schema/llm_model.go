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

type LlmModel struct {
	ent.Schema
}

func (LlmModel) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("model_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),

		field.String("model_name").
			StorageKey("model_name").Annotations(
			entgql.OrderField("MODEL_NAME")),

		field.String("provider").
			StorageKey("provider").Annotations(
			entgql.OrderField("PROVIDER")),

		field.String("model_code").
			StorageKey("model_code").Annotations(
			entgql.OrderField("MODEL_CODE")),

		field.String("api_key").
			StorageKey("api_key"),

		field.String("api_endpoint").
			StorageKey("api_endpoint").Optional().Nillable(),

		field.Bool("enabled").
			StorageKey("enabled").Default(true),

		field.Int("max_tokens").
			StorageKey("max_tokens").Default(4096),

		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),

		field.Time("updated_at").
			StorageKey("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Optional().Nillable().
			Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

func (LlmModel) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (LlmModel) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("ai", "llm_model"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
