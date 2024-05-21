package routes

import (
	"privia-staj-backend-todo/controllers"
	"privia-staj-backend-todo/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes, API endpointlerini ve ilgili controller fonksiyonlarını tanımlar.
func SetupRoutes(router *gin.Engine) {
	// Genel Rotalar
	router.POST("/api/v1/login", controllers.Login)

	// To-Do İşlemleri (Yetkilendirme Gerektirir)
	authorized := router.Group("/api/v1/todos", middlewares.AuthMiddleware())
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
