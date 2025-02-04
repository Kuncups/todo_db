package routes

import (
	"todo-api/controllers"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	taskRoutes := r.Group("/tasks")
	taskRoutes.Use(middleware.AuthMiddleware())
	{
		taskRoutes.POST("/", controllers.CreateTask)
		taskRoutes.GET("/", controllers.GetTasks)
		taskRoutes.GET("/:id", controllers.GetTaskByID)
		taskRoutes.PUT("/:id", controllers.UpdateTask)
		taskRoutes.DELETE("/:id", controllers.DeleteTask)
	}
}
