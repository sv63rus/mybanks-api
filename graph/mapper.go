package graph

import (
	"mybanks-api/ent"
	"mybanks-api/graph/model"
	"strconv"
)

func MapBank(bank *ent.Bank) *model.Bank {
	if bank == nil {
		return nil
	}

	return &model.Bank{
		ID:            strconv.Itoa(bank.ID),
		Name:          bank.Name,
		Country:       bank.Country,
		Website:       &bank.Website,
		LogoURL:       &bank.LogoURL,
		Offers:        MapOffers(bank.Edges.Offers),
		CurrencyRates: MapCurrencyRates(bank.Edges.CurrencyRates),
	}
}

func MapOffer(offer *ent.Offer) *model.Offer {
	if offer == nil {
		return nil
	}

	return &model.Offer{
		ID:          strconv.Itoa(offer.ID),
		Type:        offer.Type,
		Description: offer.Description,
		Link:        derefString(&offer.Link),
	}
}

func MapCurrencyRate(rate *ent.CurrencyRate) *model.CurrencyRate {
	if rate == nil {
		return nil
	}

	return &model.CurrencyRate{
		ID:       strconv.Itoa(rate.ID),
		Currency: string(rate.Currency),
		Rate:     rate.Rate,
	}
}

// MapOffers преобразует массив Offer из ent в model.
func MapOffers(offers []*ent.Offer) []*model.Offer {
	var result []*model.Offer
	for _, o := range offers {
		result = append(result, MapOffer(o))
	}
	return result
}

// MapCurrencyRates преобразует массив CurrencyRate из ent в model.
func MapCurrencyRates(rates []*ent.CurrencyRate) []*model.CurrencyRate {
	var result []*model.CurrencyRate
	for _, r := range rates {
		result = append(result, MapCurrencyRate(r))
	}
	return result
}

func derefString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
