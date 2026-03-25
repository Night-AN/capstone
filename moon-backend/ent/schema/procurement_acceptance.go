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

type ProcurementAcceptance struct {
	ent.Schema
}

func (ProcurementAcceptance) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_acceptance_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementAcceptanceName").
			StorageKey("procurement_acceptance_name").Annotations(entgql.OrderField("PROCUREMENT_ACCEPTANCE_NAME")),

		field.String("ProcurementAcceptanceCode").
			StorageKey("procurement_acceptance_code").Annotations(entgql.OrderField("PROCUREMENT_ACCEPTANCE_CODE")),

		field.String("ProcurementAcceptanceResult").
			StorageKey("procurement_acceptance_result").Annotations(entgql.OrderField("PROCUREMENT_ACCEPTANCE_RESULT")),

		field.String("ProcurementAcceptanceDescription").
			StorageKey("procurement_acceptance_description").Annotations(entgql.OrderField("PROCUREMENT_ACCEPTANCE_DESCRIPTION")),

		field.JSON("OtherMetadata", map[string]interface{}{}).
			StorageKey("other_metadata").Default(map[string]interface{}{}),

		field.Time("CreatedAt").
			StorageKey("created_at").Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),

		field.Time("UpdatedAt").
			StorageKey("updated_at").Optional().Nillable().Annotations(entgql.OrderField("UPDATED_AT")),

		field.UUID("procurement_implementation_id", uuid.UUID{}).StorageKey("procurement_implementation_id").Optional().Nillable(),
		field.UUID("organization_id", uuid.UUID{}).StorageKey("organization_id").Optional().Nillable(),
		field.UUID("user_id", uuid.UUID{}).StorageKey("user_id").Optional().Nillable(),
	}
}

func (ProcurementAcceptance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("procurement_implementation", ProcurementImplementation.Type).
			Ref("procurement_acceptance"). // 必须和上面一致
			Field("procurement_implementation_id").
			Unique(),
	}
}

func (ProcurementAcceptance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_acceptance"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
