package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"menu-ai-service/internal/models"
	"menu-ai-service/internal/services"
	"net/http"
)

type MenuHandler struct {
	menuService *services.MenuServiceImpl
}

func NewMenuHandler(service *services.MenuServiceImpl) *MenuHandler {
	return &MenuHandler{
		menuService: service,
	}
}

func (h *MenuHandler) Create(c *gin.Context) {
	var request models.MenuRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "ShouldBindJSON : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	created, err := h.menuService.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Menu was created with id %s", created.ID),
		"status":  http.StatusOK,
	})
}