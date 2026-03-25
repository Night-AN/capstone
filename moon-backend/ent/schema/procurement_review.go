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

type ProcurementReview struct {
	ent.Schema
}

func (ProcurementReview) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("procurement_review_id").
			Immutable().Annotations(entgql.OrderField("ID")),

		field.String("ProcurementReviewName").
			StorageKey("procurement_review_name").Annotations(entgql.OrderField("PROCUREMENT_REVIEW_NAME")),

		field.String("ProcurementReviewCode").
			StorageKey("procurement_review_code").Annotations(entgql.OrderField("PROCUREMENT_REVIEW_CODE")),

		field.String("ProcurementReviewResult").
			StorageKey("procurement_review_result").Annotations(entgql.OrderField("PROCUREMENT_REVIEW_RESULT")),

		field.String("ProcurementReviewOpinion").
			StorageKey("procurement_review_opinion").Annotations(entgql.OrderField("PROCUREMENT_REVIEW_OPINION")),

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

func (ProcurementReview) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("procurement_implementation", ProcurementImplementation.Type).
			Ref("procurement_review"). // 必须和上面一致
			Field("procurement_implementation_id").
			Unique(),

		edge.To("procurement_experts", ProcurementExpert.Type),
	}
}

func (ProcurementReview) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("biz", "procurement_review"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
