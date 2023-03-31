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

	api.POST("/user/register", controllers.RegisterUser)
	api.POST("/user/login", controllers.LoginUser)
	api.POST("/user/logout", controllers.LogoutUser)
	api.GET("/user/get/all", controllers.GetAllUser)
	api.GET("/user/get/:id", controllers.GetUserById)

	api.POST("/post/create", controllers.CreatePost)

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	if err := router.Run("localhost:3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
