package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CurrencyRate struct {
	ent.Schema
}

func (CurrencyRate) Fields() []ent.Field {
	return []ent.Field{
		field.String("currency").
			Comment("ISO 4217 currency code"),
		field.Float("rate"),
	}

}

func (CurrencyRate) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).
			Ref("currency_rates").
			Unique().
			Required(),
	}
}
