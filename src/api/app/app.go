package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApp() {
	mapUrls()

	// router from gin helps to run the app at port 8080 (client side)
	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
