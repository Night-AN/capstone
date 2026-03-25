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

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("role_id").
			Annotations(entgql.OrderField("ID")),
		field.String("role_name").
			StorageKey("role_name").
			NotEmpty().Annotations(entgql.OrderField("ROLE_NAME")),
		field.String("role_code").
			StorageKey("role_code").
			NotEmpty().Annotations(entgql.OrderField("ROLE_CODE")),
		field.String("role_description").
			StorageKey("role_description").
			NotEmpty().Annotations(entgql.OrderField("ROLE_DESCRIPTION")),
		field.String("role_flag").
			StorageKey("role_flag").
			NotEmpty().Annotations(entgql.OrderField("ROLE_FLAG")),
		field.Time("created_at").
			StorageKey("created_at").
			Default(time.Now).Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			StorageKey("updated_at").
			Optional().Annotations(entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("role"),
		edge.To("permission", Permission.Type),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "role"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
