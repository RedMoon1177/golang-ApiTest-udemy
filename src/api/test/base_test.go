package test

import (
	"golang-testing/src/api/app"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	// fmt.Println("about to start the application...")
	// TO DO FUNCTIONAL TEST -> WE HAVE TO START THE APP (with help of GIN to start the app and send the request)
	go app.StartApp()
	// fmt.Println("application started. about to start the test cases...")
	os.Exit(m.Run())
}
