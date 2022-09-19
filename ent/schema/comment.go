package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content"),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).Ref("comments").Required().Unique(),
		edge.To("author", User.Type).Required().Unique(),
		edge.To("to", User.Type).Unique(),
	}
}
