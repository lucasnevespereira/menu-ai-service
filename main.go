package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"menu-ai-service/configs"
	"menu-ai-service/internal/router"
	"menu-ai-service/internal/services"
)

func main() {
	r := gin.Default()
	config := configs.Load()
	svs := services.InitServices(config)
	router.Setup(r, svs)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
