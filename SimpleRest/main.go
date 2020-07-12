package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"pedroprado/learning-rest/model"

	"github.com/gin-gonic/gin"
)

type message struct {
	Data      string `json:"data,omitempty"`
	MessageID string `json:"messageId"`
	// Attributes struct{} `json:"attributes"`
}

type pubSubMessage struct {
	Message      message `json:"message"`
	Subscription string  `json:"subscription"`
}

var people []model.Person

func main() {

	createPeople()

	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/messages", receiveMessages)
	router.Run(":3000")
}

func getPeople(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": people})
}

func receiveMessages(context *gin.Context) {
	message := pubSubMessage{}
	bs, _ := ioutil.ReadAll(context.Request.Body)
	fmt.Printf("%s\n", bs)
	err := json.Unmarshal(bs, &message)
	fmt.Printf("%+v \n", message)
	fmt.Printf("%+v \n", err)

	context.JSON(http.StatusOK, gin.H{"status": 200, "data": "OK"})
}

func createPeople() {
	var person1 model.Person
	var person2 model.Person
	person1 = model.Person{Id: "123", Name: "Pedro", Age: 32, Address: &model.Address{Street: "José Maria de Souza", City: "São Pedro"}}
	person2 = model.Person{Id: "124", Name: "Paulo", Age: 28, Address: &model.Address{Street: "Santos Drumont", City: "São Paulo"}}

	people = append(people, person1)
	people = append(people, person2)
}
