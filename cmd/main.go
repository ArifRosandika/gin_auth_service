package main

import (
	"learn_clean_architecture/config"
	"learn_clean_architecture/internal/delivery/http/handler"
	"learn_clean_architecture/internal/delivery/http/router"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/repository"
	"learn_clean_architecture/internal/usecase"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	redisClient := config.InitRedis()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	db.AutoMigrate(&domain.User{})

	secret := viper.GetString("JWT_KEY")

	userRepo := repository.NewUserRepository(db)
	tokenRepo := repository.NewRedisTokenRepository(redisClient)
	tokenSrv := usecase.NewTokenService(secret, tokenRepo)

	userUseCase := usecase.NewUserUseCase(userRepo, tokenRepo, tokenSrv)
	authUseCase := usecase.NewAuthUseCase(tokenSrv, userRepo)

	userHandler := handler.NewUserHandler(userUseCase)
	authHandler := handler.NewAuthHandler(authUseCase)

	router.NewUserRouter(r, authHandler, userHandler, tokenSrv)

	r.Run(":8080")
}