package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang_web/Controllers"
	"golang_web/config"
	"golang_web/middleware"
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

	database.DB.AutoMigrate(&model.TodoModel{}, &model.Account{}, &model.Book{}, &model.Order{}, &model.Booklist{})

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	indexRouter := router.Group("/login")
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
	BooklistRouter := router.Group("/booklist")
	{
		BooklistRouter.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "booklist.html", nil)
		})
	}

	BookAPI := router.Group("/api/v1/book")
	{
		BookAPI.POST("", middleware.JWTAuthMiddleware(), Controllers.CreateBook)
		BookAPI.GET("", middleware.JWTAuthMiddleware(), Controllers.GetAllBooks)
		BookAPI.GET("/:id", middleware.JWTAuthMiddleware(), Controllers.GetBook)
	}

	AccountAPI := router.Group("/api/v1/account")
	{
		AccountAPI.POST("", middleware.JWTAuthMiddleware(), Controllers.CreateAccount)
		AccountAPI.GET("", middleware.JWTAuthMiddleware(), Controllers.GetAllAccount)
		AccountAPI.GET("/:email", middleware.JWTAuthMiddleware(), Controllers.GetAccount)
		AccountAPI.GET("/:email/order", middleware.JWTAuthMiddleware(), Controllers.GetOrderByAccount)
	}

	OrderAPI := router.Group("/api/v1/order")
	{
		OrderAPI.POST("", middleware.JWTAuthMiddleware(), Controllers.CreateOrder)
		OrderAPI.PATCH("/:order_id", middleware.JWTAuthMiddleware(), Controllers.ReturnOrder)
	}

	BookListAPI := router.Group("/api/v1/booklist")
	{
		BookListAPI.PATCH("", middleware.JWTAuthMiddleware(), Controllers.InsertBookInBooklist)
		BookListAPI.GET("/:account_email", middleware.JWTAuthMiddleware(), Controllers.GetBooklistByAccount)
		BookListAPI.DELETE("/:list_id/book/:book_id", middleware.JWTAuthMiddleware(), Controllers.DeleteBookFromBooklist)
	}

	LoginAPI := router.Group("/api/v1/login")
	{
		LoginAPI.POST("", Controllers.Login)
	}

	return router
}
