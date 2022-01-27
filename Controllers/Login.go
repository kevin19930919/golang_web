package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang_web/model"
	"golang_web/pkg"
	"net/http"
	"strconv"
)

// @Summary login
// @Accept  json
// @Param title body model.LoginInfo true "login"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/login [post]
func Login(context *gin.Context) {

	var logininfo model.LoginInfo
	var account model.Account

	if err := context.ShouldBindBodyWith(&logininfo, binding.JSON); err != nil {
		fmt.Println("check login query fail", err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := model.CheckAccountValid(&account, logininfo.Email, logininfo.Password); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	JwtToken, _ := jwt_pkg.GenToken(logininfo.Email, logininfo.Password)
	context.SetCookie("account_email", account.Email, 3600, "/", "", false, false)
	context.SetCookie("list_id", strconv.Itoa(account.Booklist.ID), 3600, "/", "", false, false)
	context.SetCookie("token", JwtToken, 3600, "/", "", false, false)
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{"token": JwtToken},
	})
}
