package main

import (
	"context"
	"fmt"
	"golang_web/Controllers"
	database "golang_web/config"
	"golang_web/middleware"
	"golang_web/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var ctx = context.Background()

func main() {
	server := setupServer()
	server.Run(":8080")
}

func SetupRedis() {
	pong, err := database.Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("連線redis出錯，錯誤資訊：%v", err)
	} else {
		fmt.Println("成功連線redis", pong)
	}
}

func setupServer() *gin.Engine {

	SetupRedis()

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
			c.HTML(http.StatusOK, "login.html", nil)
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
	OrderRouter := router.Group("/order")
	{
		OrderRouter.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "order.html", nil)
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
		// AccountAPI.GET("/:email/order", middleware.JWTAuthMiddleware(), Controllers.GetOrderByAccount)
	}

	OrderAPI := router.Group("/api/v1/order")
	{
		OrderAPI.POST("", middleware.JWTAuthMiddleware(), Controllers.CreateOrder)
		OrderAPI.GET("/:account_email", middleware.JWTAuthMiddleware(), Controllers.GetOrder)
		OrderAPI.PATCH("", middleware.JWTAuthMiddleware(), Controllers.ReturnOrder)
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

	TestAPI := router.Group("/api/v1/test")
	{
		TestAPI.GET("not-use-tag", Controllers.NotUseTag)
	}

	return router
}
