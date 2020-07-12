package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Get people
func GetPeople(context *gin.Context) {

	// fmt.Println(people)
	// json.NewEncoder(response).Encode(people)
	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": })
}
