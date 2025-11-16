package http

import (
	"learn_clean_architecture/internal/delivery/http/request"
	"learn_clean_architecture/internal/helper"
	"learn_clean_architecture/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	usecase usecase.UserUseCase
}

func NewUserHandler(u usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		usecase: u,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req request.RegisterUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validator := validator.New()

	if err := validator.Struct(req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecase.Register(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	helper.SuccessResponse(c, "register succcessfully", gin.H{
		"name": req.Name,
		"email": req.Email,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req request.LoginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usecase.Login(req)

	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.SuccessResponse(c, "login succcessfully", res)
}

func (h *UserHandler) Profile(c *gin.Context) {
	email, _ := c.Get("email")
	user, err := h.usecase.GetProfile(email.(string))

	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.SuccessResponse(c, "profile", gin.H{
		"name": user.Name,
		"email": strings.ToLower(user.Email),
	})
}