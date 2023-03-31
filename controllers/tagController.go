package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreateTag(c *gin.Context) {

	var tag structs.Tag

	err := c.ShouldBindJSON(&tag)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	// check title, content not empty
	if tag.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "title must be filled"})
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

	err, tag = repository.CreateTag(tag)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create tag",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully create tag",
		"data":    tag,
	})
}

func DeleteTag(c *gin.Context) {

	idTag := c.Param("id")

	var tag structs.Tag

	err := c.ShouldBindJSON(&tag)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error()})

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

	err = repository.DeleteTag(idTag)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to delete tag",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully delete tag",
	})
}

func GetAllTag(c *gin.Context) {

	var tags []structs.Tag

	tags, err := repository.GetAllTag()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to get all tag",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully get all tag",
		"data":    tags,
	})
}
