package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/response"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RegisterUser(c *gin.Context) {

	var user structs.User
	var userResponse structs.UserResponse

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error()})
		return
	}

	// check length password
	if len(user.Password) < 6 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "failed",
			"error":   "password must be more than 6 characters"})
		return
	}

	// check role must admin or user
	if user.Role != "admin" && user.Role != "user" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "failed",
			"error":   "role not valid"})
		return
	}

	err, user = repository.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error()})
		return
	}

	response.CreateUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    userResponse,
	})
}

func GetAllUser(c *gin.Context) {
	var (
		result        gin.H
		userResponses []structs.UserResponse
		userResponse  structs.UserResponse
	)

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
			"message": "You are not authorized to access this resource",
		})
		return
	}

	// Jika token valid, maka dapat melanjutkan permintaan
	err, users := repository.GetAllUser()

	// check if error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	// parsing users to userResponse
	for _, user := range users {
		userResponse = response.CreateUserResponse(user)
		userResponses = append(userResponses, userResponse)
	}

	result = gin.H{
		"success": true,
		"message": "success",
		"data":    userResponses,
	}

	c.JSON(http.StatusOK, result)
}

func GetUserById(c *gin.Context) {
	var (
		result gin.H
	)

	id := c.Param("id")

	err, user := repository.GetUserById(id)

	// check if error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"success": true,
		"message": "success",
		"data":    user,
	}

	c.JSON(http.StatusOK, result)
}

func LoginUser(c *gin.Context) {
	var (
		result gin.H
	)

	var user structs.User
	var userResponse structs.UserResponse

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error()})
		return
	}

	err, user = repository.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error()})
		return
	}

	userResponse = response.CreateUserResponse(user)

	result = gin.H{
		"success": true,
		"message": "success",
		"data":    userResponse,
	}

	c.JSON(http.StatusOK, result)
}
