package main

import "testing.rest.ginGonic/router"

var appPort = ":8080"

func main() {
	router.StartApp(appPort)
}
