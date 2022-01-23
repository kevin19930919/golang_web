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

// @Summary add account record
// @Accept  json
// @Param title body model.CreateAccountModel true "account"
// @Success 200 {string} json "{"msg":"ok"}"
// @Router /api/v1/account [post]
func CreateAccount(context *gin.Context) {
	var account model.Account
	var booklist model.Booklist
	if err := context.BindJSON(&account); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.CreatetAccountBooklist(&account, &booklist); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("success create account record")
	context.JSON(http.StatusOK, account)
}

// @Summary get all account record
// @Success 200 {string} json "{"data":[{title:title1,complete:1},{title:title2,complete:0}]}"
// @Router /api/v1/account [get]
func GetAllAccount(context *gin.Context) {
	var account []model.Account
	err := model.GetAllAccounts(&account)
	if err != nil {
		fmt.Println(err.Error())
		// context.AbortWithStatus(http.StatusNotFound)
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, account)
	}
}

// @Summary get account record
// @Param email path string true "email"
// @Success 200 {string} json "{"data":[{email:example@.xxx.com,complete:1}]}"
// @Router /api/v1/account/{email} [get]
func GetAccount(context *gin.Context) {
	type GetAccountModel struct {
		Email string `uri:"email"`
	}
	var getaccountmodel GetAccountModel
	var account model.Account

	if err := context.ShouldBindUri(getaccountmodel); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	email := context.Params.ByName("email")
	if err := model.GetAccount(&account, email); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, account)
}

// @Summary get order by account
// @Param email path string true "email"
// @Router /api/v1/account/{email}/order [get]
func GetOrderByAccount(context *gin.Context) {
	type GetAccountModel struct {
		Email string `uri:"email"`
	}
	var getaccountmodel GetAccountModel
	var account model.Account
	var order model.Order

	if err := context.ShouldBindUri(&getaccountmodel); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	email := context.Params.ByName("email")
	if err := model.GetAccount(&account, email); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := model.GetOrderByAccount(&order, &account); err != nil {
		fmt.Println(err.Error())
		context.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, account)

}

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
