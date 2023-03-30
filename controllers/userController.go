package controllers

import (
	"api-alhasanain-blog/repository"
	"api-alhasanain-blog/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {

	var user structs.User

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

	err, user = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "failed",
			"error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "success",
		"data":    user,
	})
}

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

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

	result = gin.H{
		"success": true,
		"message": "success",
		"data":    users,
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

	result = gin.H{
		"success": true,
		"message": "success",
		"data":    user,
	}

	c.JSON(http.StatusOK, result)
}
