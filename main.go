package main

import (
	"todo-api/config"
	"todo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}
