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

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("permission_id").
			Annotations(entgql.OrderField("ID")),
		field.String("permission_name").
			StorageKey("permission_name").
			NotEmpty().Annotations(entgql.OrderField("PERMISSION_NAME")),
		field.String("permission_code").
			StorageKey("permission_code").
			NotEmpty().Annotations(entgql.OrderField("PERMISSION_CODE")),
		field.String("permission_description").
			StorageKey("permission_description").
			NotEmpty().Annotations(entgql.OrderField("PERMISSION_DESCRIPTION")),
		field.String("permission_flag").
			StorageKey("permission_flag").
			NotEmpty().Annotations(entgql.OrderField("PERMISSION_FLAG")),
		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			StorageKey("updated_at").
			Optional().Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).Ref("permission"),
	}
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "permission"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
