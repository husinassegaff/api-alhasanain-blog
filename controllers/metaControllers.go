package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreateMeta(c *gin.Context) {

	var meta structs.Meta

	err := c.ShouldBindJSON(&meta)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return

	}

	// check key, content not empty
	if meta.Key == "" || meta.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "key, content must be filled"})
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

	err, meta = repository.CreateMeta(meta)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create meta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meta created",
		"data":    meta,
	})
}

func UpdateMeta(c *gin.Context) {

	var meta structs.Meta

	err := c.ShouldBindJSON(&meta)

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

	err = repository.UpdateMeta(meta)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update meta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meta updated",
	})
}

func DeleteMeta(c *gin.Context) {

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

	err = repository.DeleteMeta(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to delete meta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meta deleted",
	})
}

func GetAllMeta(c *gin.Context) {

	var metas []structs.Meta

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

	err, metas = repository.GetAllMeta()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get meta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meta fetched",
		"data":    metas,
	})
}

func GetMetaById(c *gin.Context) {

	var id = c.Param("id")

	var meta structs.Meta

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

	err, meta = repository.GetMetaById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to get meta",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Meta fetched",
		"data":    meta,
	})
}
