package router

import (
	"github.com/gin-gonic/gin"
	controller "testing.rest.ginGonic/controller"
)

var (
	router = gin.Default()
)

func StartApp(port string) {

	countryController := &controller.CountryController{}
	router.GET("/locations/countries/:country_id", countryController.GetCountry)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
