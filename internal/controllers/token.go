package controllers

import (
	"net/http"
	"test/v2/internal/auth"
	"test/v2/internal/models"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Password string `json:"password"`
}

// Token godoc
// @Summary      Generate token
// @Description  create a new token by user
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param        Token body models.User true "user name and password"
// @Success      201
// @Router       /token [post]
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	user.HashPassword(request.Password)

	// check if email exists and password is correct
	// record := database.Instance.Where("email = ?", request.Email).First(&user)
	// if record.Error != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
	// 	context.Abort()
	// 	return
	// }

	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
