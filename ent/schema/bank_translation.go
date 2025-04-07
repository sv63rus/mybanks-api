package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BankTranslation struct {
	ent.Schema
}

func (BankTranslation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("bank_id"),
		field.String("locale").NotEmpty(),
		field.String("name"),
		field.String("description"),
	}
}

func (BankTranslation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bank", Bank.Type).
			Ref("translations").
			Unique().
			Required().
			Field("bank_id"),
	}
}
