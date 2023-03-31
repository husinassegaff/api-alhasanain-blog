package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreateCategory(c *gin.Context) {

	var category structs.Category

	err := c.ShouldBindJSON(&category)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return

	}

	// check title, content not empty
	if category.Title == "" || category.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "title, content must be filled"})
		return
	}

	// Mengambil header Authorization dari permintaan
	authHeader := c.GetHeader("Authorization")

	// Memeriksa apakah header Authorization ada atau tidak
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Authorization header required",
		})
		return
	}

	// Memisahkan token dari header dengan split
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token format",
		})
		return
	}

	// Mengambil token dari hasil split
	token := splitToken[1]

	// Memverifikasi token
	err, role := repository.GetRoleWithToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "You are not admin",
		})
		return
	}

	// Membuat kategori baru
	err, category = repository.CreateCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category created",
		"data":    category,
	})

}

func UpdateCategory(c *gin.Context) {

	var category structs.Category

	err := c.ShouldBindJSON(&category)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return

	}

	// Mengambil header Authorization dari permintaan
	authHeader := c.GetHeader("Authorization")

	// Memeriksa apakah header Authorization ada atau tidak
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Authorization header required",
		})
		return
	}

	// Memisahkan token dari header dengan split
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token format",
		})
		return
	}

	// Mengambil token dari hasil split
	token := splitToken[1]

	// Memverifikasi token
	err, role := repository.GetRoleWithToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "You are not admin",
		})
		return
	}

	// Memperbarui kategori
	err = repository.UpdateCategory(category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category updated",
	})

}

func DeleteCategory(c *gin.Context) {

	var id = c.Param("id")

	// Mengambil header Authorization dari permintaan
	authHeader := c.GetHeader("Authorization")

	// Memeriksa apakah header Authorization ada atau tidak
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Authorization header required",
		})
		return
	}

	// Memisahkan token dari header dengan split
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token format",
		})
		return
	}

	// Mengambil token dari hasil split
	token := splitToken[1]

	// Memverifikasi token
	err, role := repository.GetRoleWithToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "You are not admin",
		})
		return
	}

	// check category exist or not
	err = repository.CheckCategoryById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "category not found",
		})
		return
	}

	// Menghapus kategori
	err = repository.DeleteCategory(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Category deleted",
	})

}

func GetAllCategory(c *gin.Context) {

	var categories []structs.Category

	// Mengambil semua kategori
	err, categories := repository.GetAllCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    categories,
	})
}

func GetCategoryById(c *gin.Context) {

	var id = c.Param("id")

	// Mengambil kategori berdasarkan id
	err, category := repository.GetCategoryById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    category,
	})
}
