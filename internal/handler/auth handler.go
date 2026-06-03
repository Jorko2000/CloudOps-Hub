package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/georgigeorgiev/cloudops/auth-service/internal/dto"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func (h *AuthHandler) Register(
	c *gin.Context,
) {

	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.Service.Register(
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "user registered",
		},
	)
}

func (h *AuthHandler) Login(
	c *gin.Context,
) {

	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	token, err := h.Service.Login(
		req.Email,
		req.Password,
	)

	if err != nil {

		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		dto.AuthResponse{
			AccessToken: token,
		},
	)
}

func (h *AuthHandler) Me(
	c *gin.Context,
) {

	email, _ := c.Get("email")

	c.JSON(
		http.StatusOK,
		gin.H{
			"email": email,
		},
	)
}
