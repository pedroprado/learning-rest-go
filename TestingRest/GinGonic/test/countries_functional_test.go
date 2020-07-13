package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"testing.rest.ginGonic/domain"
	"testing.rest.ginGonic/utils/errors"
)

//Here we should mock the last layers used: usually, the last layer is a
//http client interface or a database interface
//In this example, the last layer should be a http client (is is already a fake one)
func TestGetCountries(t *testing.T) {
	expectedCountry := domain.Country{
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

	response, _ := http.Get("http://localhost" + testPort + "/locations/countries/AR")
	bs, _ := ioutil.ReadAll(response.Body)
	country := domain.Country{}
	err := json.Unmarshal(bs, &country)
	if err != nil {
		fmt.Println(err.Error())
	}

	if !reflect.DeepEqual(200, response.StatusCode) {
		t.Errorf("Expected status code: %+v \n Got: %+v", 200, response.StatusCode)
	}
	if !reflect.DeepEqual(expectedCountry, country) {
		t.Errorf("Expected status code: %+v \n Got: %+v", expectedCountry, country)
	}
}

func TestGetCountriesError(t *testing.T) {
	expectedError := errors.ApiError{
		Status:  500,
		Message: "Internal Server Error",
	}

	response, _ := http.Get("http://localhost" + testPort + "/locations/countries/sss")
	bs, _ := ioutil.ReadAll(response.Body)
	apiError := errors.ApiError{}
	err := json.Unmarshal(bs, &apiError)
	if err != nil {
		fmt.Println(err.Error())
	}

	if !reflect.DeepEqual(500, response.StatusCode) {
		t.Errorf("Expected status code: %+v \n Got: %+v", 200, response.StatusCode)
	}
	if !reflect.DeepEqual(expectedError, apiError) {
		t.Errorf("Expected status code: %+v \n Got: %+v", expectedError, apiError)
	}
}
