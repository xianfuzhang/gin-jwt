package controllers

import (
	"net/http"
	"test/v2/internal/models"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary      Register user account
// @Description  create user account
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Param        User body models.User true "user name and password"
// @Success      201
// @Router       /user/register [post]
func RegisterUser(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// record := database.Instance.Create(&user)
	// if record.Error != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
	// 	context.Abort()
	// 	return
	// }
	context.JSON(http.StatusCreated, gin.H{"username": user.Name})
}
