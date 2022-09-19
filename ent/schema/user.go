package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("account").Unique(),
		field.String("password").Sensitive(),
		field.String("salt").Sensitive().Optional(),
		field.String("avatar_url").Optional(),
		field.String("intro").Optional(),
		field.Int64("coin").Default(0),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("favorite_posts", Post.Type),
		edge.To("coins", Coin.Type),
	}
}
