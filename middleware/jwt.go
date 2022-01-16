package middleware

import (
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
			context.AbortWithStatus(http.StatusNotFound)
			return
		}

		mc, err := jwt_pkg.ParseToken(parts[1])
		if err != nil {
			context.AbortWithStatus(http.StatusNotFound)
			return
		}
		context.Set("email", mc.Email)
		context.Next()
	}
}
