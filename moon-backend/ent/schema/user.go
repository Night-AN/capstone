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
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("user_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),
		field.String("Nickname").
			StorageKey("nickname").Annotations(
			entgql.OrderField("NICKNAME")),
		field.String("Fullname").
			StorageKey("full_name").Annotations(
			entgql.OrderField("FULLNAME")),
		field.String("Email").
			StorageKey("email").Annotations(
			entgql.OrderField("EMAIL")),
		field.String("PasswordHash").
			StorageKey("password_hash").Annotations(
			entgql.OrderField("PASSWORD_HASH")),
		field.Time("CreatedAt").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("UpdatedAt").
			StorageKey("updated_at").
			Optional().
			Annotations(
				entgql.OrderField("UPDATED_AT")),
		field.UUID("organization_id", uuid.UUID{}).
			StorageKey("organization_id").Optional().Nillable(),
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
