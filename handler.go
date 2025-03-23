package main

import (
	"context"
	"mybanks-api/api/gen" // путь к сгенерированному пакету
)

type handler struct{}

// ListBank реализует GET /banks
func (h *handler) ListBank(ctx context.Context, params gen.ListBankParams) (gen.ListBankRes, error) {
	banks := []gen.BankList{
		{
			ID:      1,
			Name:    "Example Bank",
			Country: "US",
			Website: gen.OptString{Set: true, Value: "https://examplebank.com"},
			LogoURL: gen.OptString{Set: true, Value: "https://examplebank.com/logo.png"},
		},
		{
			ID:      2,
			Name:    "Example Bank GE",
			Country: "GE",
			Website: gen.OptString{Set: true, Value: "https://examplebank.ge"},
			LogoURL: gen.OptString{Set: true, Value: "https://examplebank.ge/logo.png"},
		},
	}

	return (*gen.ListBankOKApplicationJSON)(&banks), nil
}

// ReadBank реализует GET /banks/{id}
func (h *handler) ReadBank(ctx context.Context, params gen.ReadBankParams) (gen.ReadBankRes, error) {
	bank := gen.BankRead{
		ID:      params.ID,
		Name:    "Example Bank111",
		Country: "US",
		Website: gen.OptString{Set: true, Value: "https://examplebank.com"},
		LogoURL: gen.OptString{Set: true, Value: "https://examplebank.ge/logo.png"},
	}

	return &bank, nil
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
