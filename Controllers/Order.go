package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"golang_web/services"
	"net/http"
	"strconv"
)

// @Summary add order record
// @Accept  json
// @Param title body service.CreateOrderByListInfo true "book-list"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/order [post]
func CreateOrder(context *gin.Context) {
	var booklist model.Booklist
	var orderinfo service.CreateOrderByListInfo

	if err := context.ShouldBindBodyWith(&orderinfo, binding.JSON); err != nil {
		fmt.Println("bind order info error", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// ======= get book-list record =======
	IntListID, _ := strconv.Atoi(orderinfo.BooklistID)
	if err := model.GetBooklistByID(&booklist, IntListID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// ======get account record ========
	var account model.Account
	if err := model.GetAccount(&account, booklist.AccountEmail); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// iter to place order
	for i := 0; i < len(orderinfo.Orders); i++ {
		// ======= get book record ========
		var book model.Book
		if err := model.GetBookByID(&book, orderinfo.Orders[i].BookID); err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		// ==========   ============
		var order model.Order
		order.EndDate = orderinfo.Orders[i].EndDate
		if err := service.CreateOderRemoveList(&order, &book, &account, &booklist); err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	fmt.Println("success create order record")
	context.JSON(http.StatusOK, booklist)
}

// // @Summary return order
// // @Param id path string true "order_id"
// // @Success 200 {string} json "{"msg":"ok"}"
// // @Router /api/v1/order/{id} [patch]
// func ReturnOrder(context *gin.Context) {
// 	var updateinfo model.ReturneOrderInfo
// 	var order model.Order

// 	if err := context.ShouldBind(&updateinfo); err != nil {
// 		fmt.Println(err.Error())
// 		context.JSON(http.StatusNotFound, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if err := service.GetOrderByID(&order, updateinfo.ID); err != nil {
// 		fmt.Println(err.Error())
// 		context.JSON(http.StatusNotFound, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if err := service.ReturnBooks(&order); err != nil {
// 		fmt.Println(err.Error())
// 		context.JSON(http.StatusNotFound, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	fmt.Println("success update order")
// 	context.JSON(http.StatusOK, order)
// }
