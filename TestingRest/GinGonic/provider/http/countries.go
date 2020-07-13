package provider

import (
	"testing.rest.ginGonic/domain"
	"testing.rest.ginGonic/utils/errors"
)

type ICountryProvider interface {
	FetchCountry(countryId string) (*domain.Country, *errors.ApiError)
}

type CountryProvider struct{}

func init() {
}

func (provider *CountryProvider) FetchCountry(countryId string) (*domain.Country, *errors.ApiError) {
	if countryId == "AR" {
		ar := &domain.Country{
			Id:     "ar_es",
			Name:   "Argentina",
			Locale: "America",
			GeoInformation: domain.GeoInformation{
				Location: domain.Location{
					Latitude:  19.222,
					Longitude: 20.000,
				},
			},
			States: []domain.State{},
		}
		return ar, nil
	}
	return nil, &errors.ApiError{
		Status:  500,
		Message: "Internal Server Error",
	}
}
