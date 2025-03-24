package dto

type CurrencyRate struct {
	ID       int     `json:"id"`
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
}

type Offer struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Link        string `json:"link,omitempty"`
}

type BankReadExtended struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Country       string         `json:"country"`
	Website       *string        `json:"website,omitempty"`
	LogoURL       *string        `json:"logo_url,omitempty"`
	Offers        []Offer        `json:"offers,omitempty"`
	CurrencyRates []CurrencyRate `json:"currency_rates,omitempty"`
}
