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

// // @Summary get all order record
// // @Success 200 {string} json "{"data":[{title:title1,complete:1},{title:title2,complete:0}]}"
// // @Router /api/v1/order [get]
// func GetAllorder(context *gin.Context) {
// 	var order []model.Order
// 	err := model.GetAllOrders(&order)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		context.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		context.JSON(http.StatusOK, order)
// 	}
// }

// // @Summary get order record by account
// // @Param title path string true "title"
// // @Success 200 {string} json "{"data":[{title:title,complete:1}]}"
// // @Router /api/v1/order [get]
// func GetOrder(context *gin.Context) {
// 	type GetOrderModel struct {
// 		title string `json:"title"`
// 	}
// 	var getordermodel GetOrderModel
// 	var order model.Order

// 	if err := context.ShouldBindUri(getordermodel); err != nil {
// 		fmt.Println(err.Error())
// 		context.AbortWithStatus(http.StatusNotFound)
// 	}
// 	title := context.Params.ByName("title")
// 	if err := model.GetOrderByAccount(&order); err != nil {
// 		fmt.Println(err.Error())
// 		context.AbortWithStatus(http.StatusNotFound)
// 	}
// 	context.JSON(http.StatusOK, order)
// }
