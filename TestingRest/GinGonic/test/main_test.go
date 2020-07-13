package test

import (
	"fmt"
	"os"
	"testing"

	"testing.rest.ginGonic/router"
)

var testPort = ":9000"

func TestMain(m *testing.M) {
	fmt.Println("About to start")
	go router.StartApp(testPort)
	os.Exit(m.Run())
}
