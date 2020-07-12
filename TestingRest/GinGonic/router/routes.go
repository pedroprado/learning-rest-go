package router

import (
	"github.com/gin-gonic/gin"
	controller "testing.rest.ginGonic/controller"
)

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}

func mapUrls() {
	router.GET("/locations/countries/:country_id", controller.GetCountry)
}
