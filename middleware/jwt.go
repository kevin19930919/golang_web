package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_web/pkg"
	"net/http"
	"strings"
)

// JWTAuthMiddleware
func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		// we want client put there jwt-token in header
		authHeader := context.Request.Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			fmt.Println("jwt token format wrong")
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		mc, err := jwt_pkg.ParseToken(parts[1])
		if err != nil {
			fmt.Println("jwt token parse error")
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		fmt.Println("jwt tokn authorize success")
		context.Set("email", mc.Email)
		context.Next()
	}
}
