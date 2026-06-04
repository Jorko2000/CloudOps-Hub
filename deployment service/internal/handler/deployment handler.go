package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"deployment-service/internal/dto"
	"deployment-service/internal/service"
)

type DeploymentHandler struct {
	Service *service.DeploymentService
}

func (h *DeploymentHandler) Deploy(
	c *gin.Context,
) {

	var req dto.DeploymentRequest

	if err := c.ShouldBindJSON(
		&req,
	); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.Service.Deploy(
		req.AppName,
		req.Image,
		req.Replicas,
	)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "deployment created",
		},
	)
}
