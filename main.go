package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang_web/Controllers"
	"golang_web/config"
	"golang_web/model"
)

func main() {
	server := setupServer()
	server.Run(":8080")
}

func setupServer() *gin.Engine {

	var err error
	var constr string
	constr = database.InitDB(database.BuildDBConfig())
	database.DB, err = gorm.Open("postgres", constr)
	if err != nil {
		fmt.Println("fail to connect database:%s", err)
	}

	database.DB.AutoMigrate(&model.TodoModel{})

	router := gin.Default()
	v1 := router.Group("/api/v1/todo")
	{
		v1.POST("", Controllers.CreateTodo)
		v1.GET("", Controllers.GetAllTodo)
		v1.GET("/:title", Controllers.GetTodo)
	}

	return router
}
