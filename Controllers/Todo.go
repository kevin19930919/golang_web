package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_web/model"
	"net/http"
)

func CreateTodo(context *gin.Context) {
	var todo model.TodoModel
	context.BindJSON(&todo)
	err := model.CreatetTodo(&todo)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println("success create todo record")
		context.JSON(http.StatusOK, todo)
	}
}
