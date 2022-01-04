package main

import (
	"github.com/gin-gonic/gin"
	"golang_web/Controllers"
)

func main() {
	server := setupServer()
	server.Run(":8080")
}

func setupServer() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/todo")
	{
		v1.POST("/", Controllers.CreateTodo)
	}

	return router
}
