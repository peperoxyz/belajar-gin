package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/book", controllers.CreateBook)
	router.GET("/book/:id", controllers.GetBook)
	router.GET("/books", controllers.GetBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	

	return router
}
