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

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("organization_id").
			Annotations(entgql.OrderField("ID")),
		field.String("organization_name").
			StorageKey("organization_name").
			NotEmpty().Annotations(entgql.OrderField("ORGANIZATION_NAME")),
		field.String("organization_code").
			StorageKey("organization_code").
			NotEmpty().Annotations(entgql.OrderField("ORGANIZATION_CODE")),
		field.String("organization_description").
			StorageKey("organization_description").
			NotEmpty().Annotations(entgql.OrderField("ORGANIZATION_DESCRIPTION")),
		field.String("organization_flag").
			StorageKey("organization_flag").
			NotEmpty().Annotations(entgql.OrderField("ORGANIZATION_FLAG")),
		field.UUID("parent_id", uuid.UUID{}).Nillable().
			StorageKey("parent_id").Annotations(entgql.OrderField("PARENET_ID")),
		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			StorageKey("updated_at").
			Optional().Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required(),
		edge.To("asset", Assets.Type),
	}
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "organization"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
