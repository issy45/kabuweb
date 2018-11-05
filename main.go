// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kabu/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
