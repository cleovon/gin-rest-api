package routes

import (
	"github.com/gin-gonic/gin"

	"gin-rest-api/src/app/controllers"
)

func RegisterRoutes(g *gin.Engine) {
	g.GET("/books", controllers.GetBooks)
	g.GET("/books/:id", controllers.GetBookByID)
	g.POST("/books", controllers.PostBook)
	g.PUT("/books/:id", controllers.PutBook)
	g.DELETE("/books/:id", controllers.DeleteBook)
}
