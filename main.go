package main

import (
	"golang-rest-server/middlewares"
	"golang-rest-server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// Initialize middlare
	router.Use(middlewares.CORSMiddleware())

	router.GET("/:collection", controllers.HandleGet)
	router.PUT("/:collection/:id", controllers.HandleUpdate)
	router.POST("/tasks", controllers.HandleAdd)
	router.DELETE("/tasks/:id", controllers.HandleDelete)

	router.Run(":8080")
}

