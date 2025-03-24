package ogent

type BankReadExtended struct {
	BankRead
	Offers        []OfferRead        `json:"offers"`
	CurrencyRates []CurrencyRateRead `json:"currency_rates"`
}

// Реализация интерфейса ReadBankRes (обязательна для ogen)
func (*BankReadExtended) ReadBankRes() {}
