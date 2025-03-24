package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Bank struct {
	ent.Schema
}

func (Bank) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.RelayConnection(), // если нужны pagination и edges
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Bank) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("country"),
		field.String("website").Optional(),
		field.String("logo_url").Optional(),
	}
}

func (Bank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("currency_rates", CurrencyRate.Type),
		edge.To("offers", Offer.Type),
	}
}
