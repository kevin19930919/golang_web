package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"net/http"
	"net/url"
)

// @Summary add order record
// @Accept  json
// @Param title body model.CreateOrderInfo true "order"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/order [post]
func CreateOrder(context *gin.Context) {
	var order model.Order
	var orderinfo model.CreateOrderInfo
	var account model.Account
	if err := context.ShouldBindBodyWith(&order, binding.JSON); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.ShouldBindBodyWith(&orderinfo, binding.JSON)
	ids := orderinfo.BookIDs

	// ======= get account record =======
	// decode url because email format got @
	email, _ := url.QueryUnescape(orderinfo.AccountEmail)
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
		if err := model.GetBookByID(&book, id); err != nil {
			fmt.Println(err.Error())
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		books = append(books, &book)
		// ======create order ============
	}
	if err := model.CreateOrder(&order, books, &account); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success create order record")
	context.JSON(http.StatusOK, order)
}

// @Summary add order record
// @Accept  json
// @Param id path string true "id"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/order/{id} [patch]
func UpdateOrder(context *gin.Context) {
	type UpdateOrderInfo struct {
		ID            int32 `json:"id"`
		OrderReturned bool  `form:"order_returned"`
		BookStatus    int32 `form:"book_status"`
	}
	var updateinfo UpdateOrderInfo
	var order model.Order

	if err := context.ShouldBind(&updateinfo); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.GetOrderByID(&order, updateinfo.ID); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.UpdateOderStatusBookReturned(&order, updateinfo.OrderReturned, updateinfo.BookStatus); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("success update order")
	context.JSON(http.StatusOK, order)
}
