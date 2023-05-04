package main

import (
	"github.com/gin-gonic/gin"
	"go-mongo/configs"
	"go-mongo/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDB()
}

func main() {
	router := gin.Default()

	routes.BookRoutes(router.Group("/book"))

	router.Run()
}
