package router

import (
	"github.com/gin-gonic/gin"
	controller "testing.rest.ginGonic/controller"
	provider "testing.rest.ginGonic/provider/http"
	"testing.rest.ginGonic/service"
)

var (
	router = gin.Default()
)

func StartApp() {

	countryController := createCountryController()
	router.GET("/locations/countries/:country_id", countryController.GetCountry)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func createCountryController() *controller.CountryController {
	provider := provider.CountryProvider{}
	service := service.CountryService{Provider: &provider}
	controller := controller.CountryController{Service: &service}

	return &controller
}
