package routes

import (
	"github.com/gin-gonic/gin"
	"go-mongo/controllers"
)

func BookRoutes(bookGroup *gin.RouterGroup) {

	bookGroup.POST("/add", controllers.CreateBook)
	bookGroup.POST("/edit/:id", controllers.UpdateBook)
	bookGroup.GET("/get/:id", controllers.GetBook)
	bookGroup.GET("/get/all", controllers.ListAllBooks)
	bookGroup.DELETE("/delete/:id", controllers.DeleteBook)
	bookGroup.DELETE("/delete", controllers.DeleteMultipleBooks)

	bookGroup.POST("/upsert", controllers.CreateUpdateBook)
	bookGroup.POST("/bulkadd", controllers.CreateManyBooks)

}
