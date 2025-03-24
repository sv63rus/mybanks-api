package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
)

type CurrencyCode string

const (
	USD CurrencyCode = "USD"
	EUR CurrencyCode = "EUR"
	GBP CurrencyCode = "GBP"
	JPY CurrencyCode = "JPY"
	CHF CurrencyCode = "CHF"
	CAD CurrencyCode = "CAD"
	AUD CurrencyCode = "AUD"
	CNY CurrencyCode = "CNY"
	RUB CurrencyCode = "RUB"
	UAH CurrencyCode = "UAH"
)

func (CurrencyCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
