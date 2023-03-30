package main

import (
	"api-alhasanain-blog/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	router := gin.Default()

	api := router.Group("/api")

	api.POST("/user/register", controllers.CreateUser)
	api.GET("/user/get/all", controllers.GetAllUser)
	api.GET("/user/get/:id", controllers.GetUserById)
	api.POST("/user/login", controllers.LoginUser)

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	if err := router.Run("localhost:3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
