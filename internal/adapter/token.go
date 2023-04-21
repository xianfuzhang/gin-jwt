package adapter

import (
	"net/http"
	"test/v2/internal/entities"
	"test/v2/internal/utils"

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
// @Param        Token body entities.User true "user name and password"
// @Success      201
// @Router       /token [post]
func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user entities.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	utils.HashPassword(request.Password)

	// check if email exists and password is correct
	// record := database.Instance.Where("email = ?", request.Email).First(&user)
	// if record.Error != nil {
	// 	context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
	// 	context.Abort()
	// 	return
	// }

	credentialError := utils.CheckPassword(user.Password, request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := utils.GenerateJWT(user.Name)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
