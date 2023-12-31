package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"menu-ai-service/internal/handlers"
	"menu-ai-service/internal/services"
)

func Setup(router *gin.Engine, services *services.Services) {

	router.Use(cors.Default())

	menuHandler := handlers.NewHandler(services)

	router.GET("/health", handlers.Health)
	router.NoRoute(handlers.NoRoute)

	router.POST("/menus", menuHandler.Save)
	router.GET("/menus/:userID", menuHandler.GetMenusByUserID)
	router.DELETE("/menus/:id", menuHandler.DeleteMenuByID)
	router.DELETE("/menus/user/:userID", menuHandler.DeleteMenusByUserID)
}
