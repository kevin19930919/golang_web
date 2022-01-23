package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"net/http"
	"strconv"
)

// @Summary create booklist record
// @Accept  json
// @Param title body model.InsertBooklistInfo true "booklist"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/booklist/{list_id}/book/{book_id} [post]
func InsertBookInBooklist(context *gin.Context) {
	var booklist model.Booklist
	var booklistinfo model.InsertBooklistInfo

	context.ShouldBindBodyWith(&booklistinfo, binding.JSON)
	id := booklistinfo.BookID
	var book model.Book

	if err := model.GetBookByID(&book, id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	IntID, _ := strconv.Atoi(booklistinfo.ID)
	if err := model.GetBooklistByID(&booklist, IntID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.InsertBookInBooklist(&booklist, &book); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success create booklist record")
	context.JSON(http.StatusOK, booklist)
}

// @Summary delete booklist record
// @Param id path string true "id"
// @Param book_id query string true "book_id"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/booklist/{list_id}/book/{book_id} [delete]
func DeleteBookFromBooklist(context *gin.Context) {
	var booklistinfo model.DeleteBooklistInfo
	var booklist model.Booklist
	var book model.Book
	if err := context.ShouldBindUri(&booklistinfo); err != nil {
		fmt.Println("request validation fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(booklistinfo.ID)

	IntID, _ := strconv.Atoi(booklistinfo.ID)

	if err := model.GetBooklistByID(&booklist, IntID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.GetBookByID(&book, booklistinfo.BookID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DeleteBookFromBooklist(&booklist, &book); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success create booklist record")
	context.JSON(http.StatusOK, booklist)

}

func GetBooklistByAccount(context *gin.Context) {
	var getbooklist model.GetBooklistInfo
	var account model.Account
	var booklist model.Booklist

	if err := context.ShouldBindUri(&getbooklist); err != nil {
		fmt.Println("get bind uri fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.GetAccount(&account, getbooklist.AccountEmail); err != nil {
		fmt.Println("get account record fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.GetBooklistByAccount(&booklist, &account); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("success get booklist record", booklist, booklist.ID)
	context.JSON(http.StatusOK, booklist)
}

func CreateOrderRemovelist(context *gin.Context) {
	var booklist model.Booklist
	var booklistinfo model.InsertBooklistInfo

	context.ShouldBindBodyWith(&booklistinfo, binding.JSON)
	id := booklistinfo.BookID
	var book model.Book

	if err := model.GetBookByID(&book, id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	IntID, _ := strconv.Atoi(booklistinfo.ID)
	if err := model.GetBooklistByID(&booklist, IntID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.InsertBookInBooklist(&booklist, &book); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success create booklist record")
	context.JSON(http.StatusOK, booklist)
}
