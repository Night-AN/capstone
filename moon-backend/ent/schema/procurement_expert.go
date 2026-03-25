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

// ProcurementExpert holds the schema definition for the ProcurementExpert entity.
type ProcurementExpert struct {
	ent.Schema
}

// Fields of the ProcurementExpert.
func (ProcurementExpert) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_expert_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementExpertName").
			StorageKey("procurement_expert_name").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_NAME")),

		field.String("ProcurementExpertCode").
			StorageKey("procurement_expert_code").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_CODE")),

		field.String("ProcurementExpertDescription").
			StorageKey("procurement_expert_description").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_DESCRIPTION")),

		field.String("ProcurementExpertFlag").
			StorageKey("procurement_expert_flag").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_FLAG")),

		field.String("ProcurementExpertJobGrade").
			StorageKey("procurement_expert_job_grade").
			Default("").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_JOB_GRADE")),

		field.String("ProcurementExpertBankName").
			StorageKey("procurement_expert_bank_name").
			Default("").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_BANK_NAME")),

		field.String("ProcurementExpertBankAccount").
			StorageKey("procurement_expert_bank_account").
			Default("").Annotations(entgql.OrderField("PROCUREMENT_EXPERT_BANK_ACCOUNT")),

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

		field.UUID("organization_id", uuid.UUID{}).
			StorageKey("organization_id").Optional().Nillable(),

		field.UUID("user_id", uuid.UUID{}).
			StorageKey("user_id").Optional().Nillable(),
	}
}

// Edges of the ProcurementExpert.
func (ProcurementExpert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("procurement_reviews", ProcurementReview.Type).
			Ref("procurement_experts"),
	}
}

// Annotations of the ProcurementExpert.
func (ProcurementExpert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_expert"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
