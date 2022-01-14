package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang_web/Controllers"
	"golang_web/config"
	"golang_web/model"
	"net/http"
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

	database.DB.AutoMigrate(&model.TodoModel{}, &model.Account{}, &model.Book{}, &model.Order{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	indexRouter := router.Group("/")
	{
		indexRouter.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	BookRouter := router.Group("/books")
	{
		BookRouter.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "books.html", nil)
		})
	}
	v1 := router.Group("/api/v1/todo")
	{
		v1.POST("", Controllers.CreateTodo)
		v1.GET("", Controllers.GetAllTodo)
		v1.GET("/:title", Controllers.GetTodo)
	}

	v2 := router.Group("/api/v1/book")
	{
		v2.POST("", Controllers.CreateBook)
		v2.GET("", Controllers.GetAllBooks)
		v2.GET("/:title", Controllers.GetBook)
	}

	v3 := router.Group("/api/v1/account")
	{
		v3.POST("", Controllers.CreateAccount)
		v3.GET("", Controllers.GetAllAccount)
		v3.GET("/:email", Controllers.GetAccount)
		v3.GET("/:email/order", Controllers.GetOrderByAccount)
	}

	v4 := router.Group("/api/v1/order")
	{
		v4.POST("", Controllers.CreateOrder)
	}

	v5 := router.Group("/api/v1/login")
	{
		v5.GET("", Controllers.Login)
	}

	return router
}
