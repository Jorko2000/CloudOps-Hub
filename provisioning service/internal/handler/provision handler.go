package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"provisioning-service/internal/dto"
	"provisioning-service/internal/service"
)

type ProvisionHandler struct {
	Service *service.ProvisionService
}

func (h *ProvisionHandler) Provision(
	c *gin.Context,
) {

	var req dto.ProvisionRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	err := h.Service.Provision(
	req.Namespace,
	req.Postgres,
	req.Redis,
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
		http.StatusOK,
		dto.ProvisionResponse{
			Status: "SUCCESS",
		},
	)
}
