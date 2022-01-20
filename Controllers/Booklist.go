package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"net/http"
	"net/url"
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
	ids := booklistinfo.BookIDs

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
	for _, id := range ids {
		var book model.Book
		if err := model.GetBook(&book, id); err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, &book)
		// ======create order ============
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
// @Param title path string true "id"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/booklist/{id} [delete]
func DeleteBooklist(context *gin.Context) {
	var booklistinfo model.DeleteBooklistInfo
	var booklist model.Booklist
	if err := context.ShouldBindUri(&booklistinfo); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.GetBooklist(&booklist, booklistinfo.ID); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.DeleteBooklist(&booklist); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success create booklist record")
	context.JSON(http.StatusOK, booklist)

}
