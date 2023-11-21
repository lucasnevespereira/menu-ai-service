package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"menu-ai-service/internal/models"
	"menu-ai-service/internal/services"
	"net/http"
)

type Handler struct {
	menuService *services.MenuServiceImpl
}

func NewHandler(service *services.Services) *Handler {
	return &Handler{
		menuService: service.MenuService,
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}

func (h *Handler) Save(c *gin.Context) {
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

func (h *Handler) GetMenusByUserID(c *gin.Context) {
	userID := c.Param("userID")
	menus, err := h.menuService.GetByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, menus)
}

func (h *Handler) DeleteMenuByID(c *gin.Context) {
	menuID := c.Param("id")
	err := h.menuService.DeleteByID(c, menuID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("menu with id %s deleted", menuID)})
}
