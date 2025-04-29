package main

import (
	"developer-notes/config"
	"developer-notes/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.NotesRoute(router)

	router.Run(":8080")

}
