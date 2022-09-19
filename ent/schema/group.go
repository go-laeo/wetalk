package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Text("intro"),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
