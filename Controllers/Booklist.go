package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"net/http"
	"net/url"
	"strconv"
)

// @Summary create booklist record
// @Accept  json
// @Param title body model.CreateBooklistInfo true "booklist"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/booklist [post]
func CreateBooklist(context *gin.Context) {
	var booklist model.Booklist
	var booklistinfo model.CreateBooklistInfo
	var account model.Account
	if err := context.ShouldBindBodyWith(&booklist, binding.JSON); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.ShouldBindBodyWith(&booklistinfo, binding.JSON)
	id := booklistinfo.BookID

	// ======= get account record =======
	// decode url because email format got @
	email, _ := url.QueryUnescape(booklistinfo.AccountEmail)
	if err := model.GetAccount(&account, email); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// ====== get all book records =======
	var books []*model.Book
	var book model.Book
	books = append(books, &book)
	if err := model.GetBookByID(&book, id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.CreateBooklist(&booklist, books, &account); err != nil {
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
	context.SetCookie("booklist_id", strconv.Itoa(booklist.ID), 3600, "/", "", false, false)
	fmt.Println("success get booklist record", booklist, booklist.ID)
	context.JSON(http.StatusOK, booklist)
}
