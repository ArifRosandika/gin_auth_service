package router

import (
	"learn_clean_architecture/internal/delivery/http/handler"
	"learn_clean_architecture/internal/delivery/http/middleware"
	"learn_clean_architecture/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, h *handler.UserHandler, tokenSrv domain.TokenService) {

	auth := r.Group("/auth")
	
	auth.POST("/refresh", h.Refresh)
	auth.POST("/login", h.Login)
	
	userRoutes := r.Group("/users") 
	
	userRoutes.POST("/register", h.Register)
	userRoutes.GET("/profile", middleware.AuthMiddleware(tokenSrv), h.Profile)
}