package http

import (

	"learn_clean_architecture/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, h *UserHandler) {

	userRoutes := r.Group("/users") 

	userRoutes.POST("/register", h.Register)
	userRoutes.POST("/login", h.Login)
	userRoutes.GET("/profile", middleware.AuthMiddleware(), h.Profile)
}