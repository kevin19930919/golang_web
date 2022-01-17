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
// @Param title body model.CreateOrder true "order"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/order [post]
func CreateOrder(context *gin.Context) {
	var order model.Order
	var book model.Book
	var account model.Account
	if err := context.ShouldBindBodyWith(&order, binding.JSON); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}

	json := make(map[string]string)
	context.ShouldBindBodyWith(&json, binding.JSON)
	id := json["book_id"]
	// decode url because email format got @
	email, _ := url.QueryUnescape(json["account_email"])

	model.GetBook(&book, id)
	model.GetAccount(&account, email)
	if err := model.CreatetOrder(&order, &book, &account); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
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
		context.AbortWithStatus(http.StatusNotFound)
	}

	if err := model.GetOrderByID(&order, updateinfo.ID); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}

	if err := model.UpdateOderStatus(&order, updateinfo.OrderReturned, updateinfo.BookStatus); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}

	fmt.Println("success update order")
	context.JSON(http.StatusOK, order)
}
