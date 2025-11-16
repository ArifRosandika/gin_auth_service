package middleware

import (
	"net/http"
	"strings"
	"learn_clean_architecture/internal/helper"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			helper.ErrorResponse(ctx, http.StatusUnauthorized, "missing authorization header")
			ctx.Abort()
			return 
		}

		tokenStr := strings.TrimPrefix(authHeader ,"Bearer ")
		email, err := helper.ValidateToken(tokenStr)

		if err != nil {
			helper.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return 
		}

		ctx.Set("email", email)

		ctx.Next()
	}
}