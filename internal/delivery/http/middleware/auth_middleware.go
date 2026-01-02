package middleware

import (
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenSrv domain.TokenService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			helper.ErrorResponse(ctx, http.StatusUnauthorized, "missing authorization header")
			ctx.Abort()
			return 
		}	

		if !strings.HasPrefix(authHeader, "Bearer") {
			helper.ErrorResponse(ctx, http.StatusUnauthorized, "invalid autorization")
			ctx.Abort()
			return 
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		email, err := tokenSrv.ValidateRefreshToken(ctx, tokenStr)

		if err != nil {
			helper.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return 
		}

		ctx.Set("email", email)
		ctx.Next()
	}
}