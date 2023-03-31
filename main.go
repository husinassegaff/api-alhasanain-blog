package main

import (
	"api-alhasanain-blog/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
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
	api.POST("/post/update", controllers.UpdatePost)
	api.POST("/post/delete/:id", controllers.DeletePost)
	api.POST("/post/update/status", controllers.ChangeStatusPost)
	api.GET("/post/get/all", controllers.GetAllPost)
	api.GET("/post/get/:id", controllers.GetPostById)

	api.POST("/category/create", controllers.CreateCategory)
	api.POST("/category/update", controllers.UpdateCategory)
	api.POST("/category/delete/:id", controllers.DeleteCategory)
	api.GET("/category/get/all", controllers.GetAllCategory)
	api.GET("/category/get/:id", controllers.GetCategoryById)

	api.POST("/meta/create", controllers.CreateMeta)
	api.POST("/meta/update", controllers.UpdateMeta)
	api.POST("/meta/delete/:id", controllers.DeleteMeta)
	api.GET("/meta/get/all", controllers.GetAllMeta)
	api.GET("/meta/get/:id", controllers.GetMetaById)

	api.POST("/tag/create", controllers.CreateTag)
	api.POST("/tag/delete/:id", controllers.DeleteTag)
	api.GET("/tag/get/all", controllers.GetAllTag)

	api.POST("/comment/create", controllers.CreateComment)
	api.POST("/comment/delete/:id", controllers.DeleteComment)
	api.GET("/comment/get/all", controllers.GetAllComment)

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	err := godotenv.Load("config/.env") // Load the .env file
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err.Error())
		os.Exit(1)
	}

	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
