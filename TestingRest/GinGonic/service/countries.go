package service

import (
	"testing.rest.ginGonic/domain"
	provider "testing.rest.ginGonic/provider/http"
	"testing.rest.ginGonic/utils/errors"
)

type CountryServiceInterface interface {
	GetCountry(countryId string) (*domain.Country, *errors.ApiError)
}

type CountryService struct {
	Provider provider.CountryProviderInterface
}

func (service *CountryService) GetCountry(countryId string) (*domain.Country, *errors.ApiError) {

	return service.Provider.FetchCountry(countryId)
}
