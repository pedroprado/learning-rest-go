package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	"testing.rest.ginGonic/domain"
	"testing.rest.ginGonic/utils/errors"
)

//For testing the Controller Unitary, we should mock the Layer Service!
//For Mocking the Layer Service, we need to Implement is as an INTERFACE
type mockCountryService struct {
	mockGetFunc func(countryId string) (*domain.Country, *errors.ApiError)
}

func (mock *mockCountryService) GetCountry(countryId string) (*domain.Country, *errors.ApiError) {
	return mock.mockGetFunc(countryId)
}

var mockService mockCountryService
var controller *CountryController

func TestMain(m *testing.M) {
	mockService = mockCountryService{
		mockGetFunc: func(countryId string) (*domain.Country, *errors.ApiError) {
			//case valid country_id
			if countryId == "AR" {
				return &domain.Country{
					Id:     "ar_es",
					Name:   "Argentina",
					Locale: "America",
				}, nil
			}
			//case invalid country_id
			return nil, &errors.ApiError{
				Status:  500,
				Message: "Internal Server Error",
			}
		},
	}
	CountryService = &mockService
	controller = &CountryController{}

	os.Exit(m.Run())
}

func TestGetCountry(t *testing.T) {
	expectedStatus := 200
	expectedCountry := domain.Country{
		Id:     "ar_es",
		Name:   "Argentina",
		Locale: "America",
	}

	response := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(response)
	testContext.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	testContext.Params = gin.Params{{Key: "country_id", Value: "AR"}}

	controller.GetCountry(testContext)

	country := domain.Country{}
	err := json.Unmarshal(response.Body.Bytes(), &country)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	if !reflect.DeepEqual(expectedStatus, response.Code) {
		t.Errorf("Expected status: %v\n Got: %v\n", expectedStatus, response.Code)
	}
	if !reflect.DeepEqual(expectedCountry, country) {
		t.Errorf("Expected country: %v\n Got: %v\n", expectedCountry, country)
	}

}

func TestGetCountryError(t *testing.T) {
	expectedStatus := 500
	expectedError := errors.ApiError{
		Status:  500,
		Message: "Internal Server Error",
	}

	response := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(response)
	testContext.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	testContext.Params = gin.Params{{Key: "country_id", Value: "invalid"}}

	controller.GetCountry(testContext)

	apiErr := errors.ApiError{}
	err := json.Unmarshal(response.Body.Bytes(), &apiErr)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
	}

	if !reflect.DeepEqual(expectedStatus, response.Code) {
		t.Errorf("Expected status: %v\n Got: %v\n", expectedStatus, response.Code)
	}
	if !reflect.DeepEqual(expectedError, apiErr) {
		t.Errorf("Expected error: %v\n Got: %v\n", expectedError, apiErr)
	}
}
