package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content"),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique().Required(),
		edge.From("group", Group.Type).Ref("posts").Unique(),
		edge.From("favorite_users", User.Type).Ref("favorite_posts"),
		edge.To("comments", Comment.Type),
	}
}
