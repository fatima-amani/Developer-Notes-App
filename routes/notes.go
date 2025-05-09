package routes

import (
	"developer-notes/controller"

	"github.com/gin-gonic/gin"
)

func NotesRoute(router *gin.Engine) {
	router.GET("/", controller.GetNotes)
	router.GET("/:id", controller.GetNote)
	router.POST("/", controller.CreateNote)
	router.PUT("/:id", controller.UpdateNote)
	router.DELETE("/:id", controller.DeleteNote)
}
