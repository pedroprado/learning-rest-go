package controller

import (
	"github.com/gin-gonic/gin"
	"testing.rest.ginGonic/service"
)

type CountryControllerInterface interface {
	GetCountry(c *gin.Context)
}

type CountryController struct {
	Service service.CountryServiceInterface
}

func (controller *CountryController) GetCountry(c *gin.Context) {
	country, err := controller.Service.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(200, country)
}
