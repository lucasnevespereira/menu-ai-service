package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"menu-ai-service/internal/handlers"
	"menu-ai-service/internal/services"
)

func Setup(router *gin.Engine, services *services.Services) {

	router.Use(cors.Default())

	_ = handlers.NewHandler(services)

	router.GET("/health", handlers.Health)

	router.NoRoute(handlers.NoRoute)
}
