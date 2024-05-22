package main

import (
	"privia-staj-backend-todo/middlewares"
	"privia-staj-backend-todo/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)               // Rotaları ayarla
	router.Use(middlewares.CORSMiddleware()) // CORS middleware'ı ekle
	router.Run(":7878")
}
