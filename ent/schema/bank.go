package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Bank struct {
	ent.Schema
}

func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("country"),
		field.String("website").Optional(),
		field.String("logo_url").Optional(),
		field.String("test").Optional(),
	}
}

func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency_rates", CurrencyRate.Type),
		edge.To("offers", Offer.Type),
	}
}
