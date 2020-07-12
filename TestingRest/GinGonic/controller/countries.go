package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"testing.rest.ginGonic/service"
)

func GetCountry(c *gin.Context) {

	country, err := service.GetCountry(c.Param("country_id"))
	if err != nil {
		fmt.Printf("%+v", err)
		c.JSON(err.Status, err)
		return
	}
	c.JSON(200, country)
}
