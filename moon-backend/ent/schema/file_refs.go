package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FileRef holds the schema definition for the FileRef entity.
type FileRef struct {
	ent.Schema
}

// Fields of the FileRef.
func (FileRef) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("file_ref_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),
		field.UUID("file_id", uuid.UUID{}).
			StorageKey("file_id").Optional(),
	}
}

// Edges of the FileRef.
func (FileRef) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("file", File.Type),
	}
}

// Annotations of the FileRef.
func (FileRef) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "file_refs"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
