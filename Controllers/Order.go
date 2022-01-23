package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"net/http"
	"strconv"
)

// @Summary add order record
// @Accept  json
// @Param title body model.CreateOrderByListInfo true "order"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/order [post]
func CreateOrder(context *gin.Context) {
	var order model.Order
	var booklist model.Booklist
	var orderinfo model.CreateOrderByListInfo

	if err := context.ShouldBindBodyWith(&orderinfo, binding.JSON); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	// ======= get book list record =======
	IntListID, _ := strconv.Atoi(orderinfo.BooklistID)
	if err := model.GetBooklistByID(&booklist, IntListID); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.CreateOderRemoveList(&order, &booklist); err != nil {
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
