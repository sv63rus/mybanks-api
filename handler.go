package main

import (
	"context"
	"fmt"
	"mybanks-api/ent"
	"mybanks-api/ent/bank"
	"mybanks-api/ent/ogent"     // путь к сгенерированному пакету
	gen "mybanks-api/ent/ogent" // путь к сгенерированному пакету
)

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func deref(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

// ListBank реализует GET /banks
func (h *handler) ListBank(ctx context.Context, params gen.ListBankParams) (gen.ListBankRes, error) {
	data, err := h.client.Bank.
		Query().
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying data: %w", err)
	}

	var result []gen.BankList
	for _, datum := range data {
		result = append(result, gen.BankList{
			ID:      int(datum.ID),
			Name:    datum.Name,
			Country: datum.Country,
			Website: gen.OptString{Set: datum.Website != "", Value: deref(&datum.Website)},
			LogoURL: gen.OptString{Set: datum.LogoURL != "", Value: deref(&datum.LogoURL)},
		})
	}

	return (*gen.ListBankOKApplicationJSON)(&result), nil
}

// ReadBank реализует GET /banks/{id}
func (h *handler) ReadBank(ctx context.Context, params gen.ReadBankParams) (gen.ReadBankRes, error) {

	bankDetail, err := h.client.Bank.Query().Where(bank.ID(params.ID)).WithOffers().WithCurrencyRates().Only(ctx)

	if nil != err {
		return nil, err
	}

	var offers []gen.OfferRead
	for _, o := range bankDetail.Edges.Offers {
		offers = append(offers, gen.OfferRead{
			ID:          int(o.ID),
			Type:        o.Type,
			Description: o.Description,
			Link:        gen.OptString{Set: o.Link != "", Value: deref(&o.Link)}.Value,
		})
	}

	var rates []gen.CurrencyRateRead
	for _, r := range bankDetail.Edges.CurrencyRates {
		rates = append(rates, gen.CurrencyRateRead{
			ID:       int(r.ID),
			Currency: r.Currency,
			Rate:     r.Rate,
		})
	}

	return &gen.BankReadExtended{
		ID:            int(bankDetail.ID),
		Name:          bankDetail.Name,
		Country:       bankDetail.Country,
		Website:       gen.OptString{Set: bankDetail.Website != "", Value: deref(&bankDetail.Website)},
		LogoURL:       gen.OptString{Set: bankDetail.LogoURL != "", Value: deref(&bankDetail.LogoURL)},
		Offers:        offers,
		CurrencyRates: rates,
	}, nil
}

func (h *handler) CreateBank(ctx context.Context, params *gen.CreateBankReq) (gen.CreateBankRes, error) {
	// TODO: implement CreateBank
	var result gen.CreateBankRes
	return result, nil
}

func (h *handler) CreateCurrencyRate(ctx context.Context, params *gen.CreateCurrencyRateReq) (gen.CreateCurrencyRateRes, error) {
	// TODO: implement CreateCurrencyRate
	var result gen.CreateCurrencyRateRes
	return result, nil
}

func (h *handler) CreateOffer(ctx context.Context, params *gen.CreateOfferReq) (gen.CreateOfferRes, error) {
	// TODO: implement CreateOffer
	var result gen.CreateOfferRes
	return result, nil
}

func (h *handler) DeleteBank(ctx context.Context, params gen.DeleteBankParams) (gen.DeleteBankRes, error) {
	// TODO: implement DeleteBank
	var result gen.DeleteBankRes
	return result, nil
}

func (h *handler) DeleteCurrencyRate(ctx context.Context, params gen.DeleteCurrencyRateParams) (gen.DeleteCurrencyRateRes, error) {
	// TODO: implement DeleteCurrencyRate
	var result gen.DeleteCurrencyRateRes
	return result, nil
}

func (h *handler) DeleteOffer(ctx context.Context, params gen.DeleteOfferParams) (gen.DeleteOfferRes, error) {
	// TODO: implement DeleteOffer
	var result gen.DeleteOfferRes
	return result, nil
}

func (h *handler) ListBankCurrencyRates(ctx context.Context, params gen.ListBankCurrencyRatesParams) (gen.ListBankCurrencyRatesRes, error) {
	// TODO: implement ListBankCurrencyRates
	var result gen.ListBankCurrencyRatesRes
	return result, nil
}

func (h *handler) ListBankOffers(ctx context.Context, params gen.ListBankOffersParams) (gen.ListBankOffersRes, error) {
	// TODO: implement ListBankOffers
	var result gen.ListBankOffersRes
	return result, nil
}

func (h *handler) ListCurrencyRate(ctx context.Context, params gen.ListCurrencyRateParams) (gen.ListCurrencyRateRes, error) {
	// TODO: implement ListCurrencyRate
	var result gen.ListCurrencyRateRes
	return result, nil
}

func (h *handler) ListOffer(ctx context.Context, params gen.ListOfferParams) (gen.ListOfferRes, error) {
	// TODO: implement ListOffer
	var result gen.ListOfferRes
	return result, nil
}

func (h *handler) ReadCurrencyRate(ctx context.Context, params gen.ReadCurrencyRateParams) (gen.ReadCurrencyRateRes, error) {
	// TODO: implement ReadCurrencyRate
	var result gen.ReadCurrencyRateRes
	return result, nil
}

func (h *handler) ReadCurrencyRateBank(ctx context.Context, params gen.ReadCurrencyRateBankParams) (gen.ReadCurrencyRateBankRes, error) {
	// TODO: implement ReadCurrencyRateBank
	var result gen.ReadCurrencyRateBankRes
	return result, nil
}

func (h *handler) ReadOffer(ctx context.Context, params gen.ReadOfferParams) (gen.ReadOfferRes, error) {
	// TODO: implement ReadOffer
	var result gen.ReadOfferRes
	return result, nil
}

func (h *handler) ReadOfferBank(ctx context.Context, params gen.ReadOfferBankParams) (gen.ReadOfferBankRes, error) {
	// TODO: implement ReadOfferBank
	var result gen.ReadOfferBankRes
	return result, nil
}

func (h *handler) UpdateBank(ctx context.Context, req *gen.UpdateBankReq, params gen.UpdateBankParams) (gen.UpdateBankRes, error) {
	// TODO: implement UpdateBank
	var result gen.UpdateBankRes
	return result, nil
}

func (h *handler) UpdateCurrencyRate(ctx context.Context, req *gen.UpdateCurrencyRateReq, params gen.UpdateCurrencyRateParams) (gen.UpdateCurrencyRateRes, error) {
	// TODO: implement UpdateCurrencyRate
	var result gen.UpdateCurrencyRateRes
	return result, nil
}

func (h *handler) UpdateOffer(ctx context.Context, req *gen.UpdateOfferReq, params gen.UpdateOfferParams) (gen.UpdateOfferRes, error) {
	// TODO: implement UpdateOffer
	var result gen.UpdateOfferRes
	return result, nil
}
