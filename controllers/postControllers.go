package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CreatePost(c *gin.Context) {

	var post structs.Post

	err := c.ShouldBindJSON(&post)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		return

	}

	// check title, content, status not empty
	if post.Title == "" || post.Content == "" || post.Status == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "title, content, status must be filled"})

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

	err, post = repository.CreatePost(post)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": post})

}

func UpdatePost(c *gin.Context) {

	var post structs.Post

	err := c.ShouldBindJSON(&post)

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

	// check post exist or not
	isExist := repository.CheckPostById(post.ID)

	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	err = repository.UpdatePost(post)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success"})

}

func DeletePost(c *gin.Context) {

	id := c.Param("id")

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

	// check post exist or not
	isExist := repository.CheckPostById(id)

	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	err = repository.DeletePost(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success"})

}

func ChangeStatusPost(c *gin.Context) {

	var post structs.Post

	err := c.ShouldBindJSON(&post)

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

	// check post exist or not
	isExist := repository.CheckPostById(post.ID)

	if isExist == false {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	err = repository.ChangeStatusPost(post.ID, post.Status)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success"})

}

func GetAllPost(c *gin.Context) {

	err, posts := repository.GetAllPost()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    posts})

}

func GetPostById(c *gin.Context) {

	id := c.Param("id")

	err, post := repository.GetPostById(id)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    post})

}
