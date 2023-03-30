package main

import (
	"api-alhasanain-blog/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	router := gin.Default()

	// User
	router.POST("/register", controllers.CreateUser)
	router.GET("/user", controllers.GetAllUser)

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	if err := router.Run("localhost:3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
