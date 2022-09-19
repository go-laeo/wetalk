package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Coin holds the schema definition for the Coin entity.
type Coin struct {
	ent.Schema
}

// Fields of the Coin.
func (Coin) Fields() []ent.Field {
	return []ent.Field{
		field.String("deal"),
		field.Int64("amount"),
		field.Int64("balance"),
		field.Time("created_at").Default(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Annotations(entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
	}
}

// Edges of the Coin.
func (Coin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("coins").Required().Unique(),
	}
}
