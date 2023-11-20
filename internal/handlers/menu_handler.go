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

func (h *MenuHandler) Save(c *gin.Context) {
	var request models.MenuSaveRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "ShouldBindJSON : " + err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	created, err := h.menuService.Save(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  err.Error(),
			"status": http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Menu was saved with id %s", created.ID),
		"status":  http.StatusOK,
	})
}

func (h *MenuHandler) GetMenusByUserID(c *gin.Context) {
	userID := c.Param("userID")
	menus, err := h.menuService.GetByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, menus)
}
