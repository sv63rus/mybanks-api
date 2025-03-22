package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Offer struct {
	ent.Schema
}

func (Offer) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.String("description"),
		field.String("link"),
	}
}

func (Offer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).
			Ref("offers").
			Unique().
			Required(),
	}
}
