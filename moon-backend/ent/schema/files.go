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

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).
			StorageKey("file_id").
			Immutable().Annotations(
			entgql.OrderField("ID")),
		field.String("FileName").
			StorageKey("file_name").Annotations(
			entgql.OrderField("FILE_NAME")),
		field.String("FileType").
			StorageKey("file_type").Annotations(
			entgql.OrderField("FILE_TYPE")),
		field.Int64("FileSize").
			StorageKey("file_size").Annotations(
			entgql.OrderField("FILE_SIZE")),
		field.Time("CreatedAt").
			StorageKey("created_at").
			Default(time.Now).
			Annotations(entgql.OrderField("CREATED_AT")),
		field.Time("UpdatedAt").
			StorageKey("updated_at").
			Optional().
			Annotations(
				entgql.OrderField("UPDATED_AT")),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file_refs", FileRef.Type).Ref("file").Required().Unique(),
	}
}

// Annotations of the File.
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.SchemaTable("systems", "files"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
