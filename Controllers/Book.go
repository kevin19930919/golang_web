package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_web/model"
	"net/http"
)

// @Summary add book record
// @Accept  json
// @Param title body model.CreateBookModel true "book"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/book [post]
func CreateBook(context *gin.Context) {

	var book model.Book
	if err := context.BindJSON(&book); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.CreatetBook(&book); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("success create book record")
	context.JSON(http.StatusOK, book)
}

// @Summary get all book record
// @Success 200 {string} json "{"data":[{title:title1,id:0},{title:title2,id:1}]}"
// @Router /api/v1/book [get]
func GetAllBooks(context *gin.Context) {
	var book []model.Book
	err := model.GetAllBooks(&book)
	if err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, book)
	}
}

// @Summary get book record
// @Param id path string true "id"
// @Success 200 {string} json "{"data":[{title:title,complete:1}]}"
// @Router /api/v1/book/{id} [get]
func GetBook(context *gin.Context) {
	type GetBookModel struct {
		ID string `json:"id"`
	}
	var getbookmodel GetBookModel
	var book model.Book

	if err := context.ShouldBindUri(getbookmodel); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := context.Params.ByName("id")
	if err := model.GetBookByID(&book, id); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, book)
}
