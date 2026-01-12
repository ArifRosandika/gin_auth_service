package router

import (
	"learn_clean_architecture/internal/delivery/http/handler"
	"learn_clean_architecture/internal/delivery/http/middleware"
	"learn_clean_architecture/internal/domain"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, a *handler.AuthHandler, h *handler.UserHandler, tokenSrv domain.TokenService) {

	auth := r.Group("/auth")
	
	auth.POST("/refresh", a.Refresh)
	auth.POST("/login", a.Login)
	
	userRoutes := r.Group("/users") 
	
	userRoutes.POST("/register", h.Register)
	userRoutes.GET("/profile", middleware.AuthMiddleware(tokenSrv), h.Profile)
	userRoutes.DELETE("/logout", h.Logout)
}