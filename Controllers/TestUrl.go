package Controllers

import (
	service "golang_web/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary check what happen if not use json tag
// @Router /api/v1/test/not-use-tag [get]
func NotUseTag(context *gin.Context) {
	type Response struct {
		TestOut service.TestIn `json:"test_out"`
	}

	inner := service.TestIn{
		Name: "kevin",
		Age:  28,
	}

	res := Response{
		TestOut: inner,
	}

	context.JSON(http.StatusOK, res)

}
