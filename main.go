package main

import (
	"developer-notes/config"

	"github.com/gin-gonic/gin"
	"honnef.co/go/tools/config"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)

	router.Run(":8080")

}
