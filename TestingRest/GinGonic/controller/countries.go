package controller

import (
	"github.com/gin-gonic/gin"
	"testing.rest.ginGonic/service"
)

var CountryService service.CountryServiceInterface

type CountryControllerInterface interface {
	GetCountry(c *gin.Context)
}

type CountryController struct {
}

func init() {
	CountryService = &service.CountryService{}
}

func (controller *CountryController) GetCountry(c *gin.Context) {
	country, err := CountryService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(200, country)
}
