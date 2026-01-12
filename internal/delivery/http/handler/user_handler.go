package handler

import (
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	usecase domain.UserUseCase
}

func NewUserHandler(u domain.UserUseCase) *UserHandler {
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

	if err := h.usecase.Register(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	helper.CreatedResponse(c, "register succcessfully", gin.H{
		"name": req.Name,
		"email": req.Email,
	})
}

func (h *UserHandler) Profile(c *gin.Context) {
	userID := c.GetUint("user_id")
	if userID == 0 {
		helper.ErrorResponse(c, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := h.usecase.GetProfile(c.Request.Context(), userID)

	if err != nil {
		helper.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	helper.SuccessResponse(c, "user profile", gin.H{
		"name": user.Name,
		"user_id": user.ID,
		"email": strings.ToLower(user.Email),
	})
}

func (h *UserHandler) Logout(c *gin.Context) {
	var req request.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err  := h.usecase.Logout(c.Request.Context() , req.Token); err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.SuccessResponse(c, "logout succcessfully", nil)
}