package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/semihbugrasezer/privia-todo-api/controllers"
	"github.com/semihbugrasezer/privia-todo-api/middlewares"
)

// SetupRoutes, API endpointlerini ve ilgili controller fonksiyonlarını tanımlar.
func SetupRoutes(router *gin.Engine) {
	// Genel Rotalar
	router.POST("/api/v1/login", controllers.Login)

	// To-Do İşlemleri (Yetkilendirme Gerektirir)
	authorized := router.Group("/api/v1/todos")
	authorized.Use(middlewares.AuthMiddleware()) // Yetkilendirme middleware'ı
	{
		authorized.GET("/", controllers.GetTodos)
		authorized.POST("/", controllers.CreateTodoList)
		authorized.GET("/:id", controllers.GetTodo)
		authorized.PUT("/:id", controllers.UpdateTodoList)
		authorized.DELETE("/:id", controllers.DeleteTodoList)

		// To-Do Öğesi İşlemleri (Yetkilendirme Gerektirir)
		items := authorized.Group("/:todoId/items")
		{
			items.GET("/", controllers.GetItems)
			items.POST("/", controllers.CreateTodoItem)
			items.GET("/:itemId", controllers.GetItem)
			items.PUT("/:itemId", controllers.UpdateTodoItem)
			items.DELETE("/:itemId", controllers.DeleteTodoItem)
		}
	}
}
