package main

import (
	"github.com/gin-gonic/gin"
	"github.com/semihbugrasezer/privia-todo-api/mock"
	"github.com/semihbugrasezer/privia-todo-api/controllers"
    "github.com/semihbugrasezer/privia-todo-api/middlewares"

)

func main() {
	mock.LoadMockData() // Mock verileri yükle
	router := gin.Default()
	routes.SetupRoutes(router)               // Rotaları ayarla
	router.Use(middlewares.CORSMiddleware()) // CORS middleware'ı ekle
	router.Run(":7878")
}
