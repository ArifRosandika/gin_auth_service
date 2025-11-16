package main

import (
	"learn_clean_architecture/config"
	"learn_clean_architecture/internal/entity"
	"learn_clean_architecture/internal/repository"
	"learn_clean_architecture/internal/usecase"
	"learn_clean_architecture/internal/delivery/http"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	r := gin.Default()

	db.AutoMigrate(&entity.User{})

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := http.NewUserHandler(userUseCase)

	http.NewUserRouter(r, userHandler)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Run(":8080")
}