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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).StorageKey("user_id").
			Default(uuid.New).Immutable().
			Annotations(entgql.OrderField("ID")),
		field.String("nickname").NotEmpty().
			Annotations(entgql.OrderField("NICKNAME")),
		field.String("fullname").NotEmpty().
			Annotations(entgql.OrderField("FULLNAME")),
		field.String("email").NotEmpty().
			Annotations(entgql.OrderField("EMAIL")),
		field.String("password_hash").NotEmpty().
			Annotations(entgql.OrderField("PASSWORD_HASH")),
		field.Time("created_at").
			Default(time.Now).Immutable().
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("updated_at").
			Default(time.Now).Immutable().
			UpdateDefault(time.Now()).
			Annotations(entgql.OrderField("UPDATED_AT")),
		field.UUID("organization_id", uuid.UUID{}).Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("organization", Organization.Type).Field("organization_id").Ref("user").Unique(),
		edge.To("role", Role.Type),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "user"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
