package routes

import (
	"privia-staj-backend-todo/controllers"
	"privia-staj-backend-todo/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up API endpoints and their corresponding controller functions
func SetupRoutes(router *gin.Engine) {
	// CORS Middleware
	router.Use(middlewares.CORSMiddleware())

	// General Routes
	router.POST("/api/v1/login", controllers.Login)

	// Todo Operations (Authorization Required)
	authorized := router.Group("/api/v1/todos", middlewares.AuthMiddleware())
	{
		authorized.GET("/", controllers.GetTodos)
		authorized.POST("/", controllers.CreateTodoList)
		authorized.GET("/:todoId", controllers.GetTodo)
		authorized.PUT("/:todoId", controllers.UpdateTodoList)
		authorized.DELETE("/:todoId", controllers.DeleteTodoList)

		// Todo Item Operations (Authorization Required)
		items := authorized.Group("/:todoId/items")
		{
			items.GET("/", controllers.GetItems)
			items.POST("/", controllers.CreateTodoItem)
			items.GET("/:itemId", controllers.GetItems)
			items.PUT("/:itemId", controllers.UpdateTodoItem)
			items.DELETE("/:itemId", controllers.DeleteTodoItem)
		}
	}
}
