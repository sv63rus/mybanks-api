package mapper

import (
	"mybanks-api/ent"
	"mybanks-api/internal/dto"
)

func MapBankToDTO(bank *ent.Bank) dto.BankReadExtended {
	var offers []dto.Offer
	for _, o := range bank.Edges.Offers {
		offers = append(offers, dto.Offer{
			ID:          int(o.ID),
			Type:        o.Type,
			Description: o.Description,
			Link:        o.Link,
		})
	}

	var rates []dto.CurrencyRate
	for _, r := range bank.Edges.CurrencyRates {
		rates = append(rates, dto.CurrencyRate{
			ID:       int(r.ID),
			Currency: r.Currency,
			Rate:     r.Rate,
		})
	}

	return dto.BankReadExtended{
		ID:            int(bank.ID),
		Name:          bank.Name,
		Country:       bank.Country,
		Website:       optionalString(bank.Website),
		LogoURL:       optionalString(bank.LogoURL),
		Offers:        offers,
		CurrencyRates: rates,
	}
}

func optionalString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
