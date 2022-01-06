package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_web/model"
	"net/http"
)

// @Summary add todo record
// @Accept  json
// @Param title body model.CreateTodoModel true "todo"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/todo [post]
func CreateTodo(context *gin.Context) {
	var todo model.TodoModel
	if err := context.BindJSON(&todo); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}

	if err := model.CreatetTodo(&todo); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}
	fmt.Println("success create todo record")
	context.JSON(http.StatusOK, todo)
}

// @Summary get all todo record
// @Success 200 {string} json "{"data":[{title:title1,complete:1},{title:title2,complete:0}]}"
// @Router /api/v1/todo [get]
func GetAllTodo(context *gin.Context) {
	var todo []model.TodoModel
	err := model.GetAllTodos(&todo)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

// @Summary get todo record
// @Param title path string true "title"
// @Success 200 {string} json "{"data":[{title:title,complete:1}]}"
// @Router /api/v1/todo/{title} [get]
func GetTodo(context *gin.Context) {
	type GetTodoModel struct {
		title string `json:"title"`
	}
	var gettodomodel GetTodoModel
	var todo model.TodoModel

	if err := context.ShouldBindUri(gettodomodel); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}
	title := context.Params.ByName("title")
	if err := model.GetTodo(&todo, title); err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	}
	context.JSON(http.StatusOK, todo)
}
