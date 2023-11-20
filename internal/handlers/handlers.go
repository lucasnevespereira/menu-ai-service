package handlers

import (
	"github.com/gin-gonic/gin"
	"menu-ai-service/internal/services"
	"net/http"
)

type Handler struct {
	menuService *services.MenuServiceImpl
}

func NewHandler(services services.Services) *Handler {
	return &Handler{
		menuService: services.MenuService,
	}
}

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "up"})
}

func NoRoute(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}
