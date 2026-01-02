package handler

import (
	"learn_clean_architecture/internal/delivery/http/dto/request"
	"learn_clean_architecture/internal/domain"
	"learn_clean_architecture/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	usecase domain.AuthUseCase
}

func NewAuthHandler(u domain.AuthUseCase) *AuthHandler {
	return &AuthHandler{
		usecase: u,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req request.LoginUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usecase.Login(c.Request.Context(), req)

	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.SuccessResponse(c, "login succcessfully", res)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	var req request.AuthRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.usecase.Refresh(c.Request.Context(), req.RefreshToken) 

	if err != nil {
		helper.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	helper.SuccessResponse(c, "refresh token", res)
}