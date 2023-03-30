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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// check length password
	if len(user.Password) < 6 {
		c.JSON(400, gin.H{
			"message": "failed",
			"error":   "Password must be more than 6 characters"})
		return
	}

	// check role must admin or user
	if user.Role != "admin" && user.Role != "user" {
		c.JSON(400, gin.H{
			"message": "failed",
			"error":   "Role not valid"})
		return
	}

	err, user = repository.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
			"error":   err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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
			"message": "failed",
			"error":   err.Error(),
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	result = gin.H{
		"message": "success",
		"data":    users,
	}

	c.JSON(http.StatusOK, result)
}
