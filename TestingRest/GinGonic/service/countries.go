package service

import (
	"testing.rest.ginGonic/domain"
	provider "testing.rest.ginGonic/provider/http"
	"testing.rest.ginGonic/utils/errors"
)

var CountryProvider provider.CountryProviderInterface

type CountryServiceInterface interface {
	GetCountry(countryId string) (*domain.Country, *errors.ApiError)
}

type CountryService struct {
}

func init() {
	CountryProvider = &provider.CountryProvider{}
}

func (service *CountryService) GetCountry(countryId string) (*domain.Country, *errors.ApiError) {

	return CountryProvider.FetchCountry(countryId)
}
