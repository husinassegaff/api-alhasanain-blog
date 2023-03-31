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

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	if err := router.Run("localhost:3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
