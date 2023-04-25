package adapter

import (
	"net/http"
	"test/v2/internal/adapter/repository"
	"test/v2/internal/entities"
	"test/v2/internal/service"
	"test/v2/internal/utils"

	"github.com/gin-gonic/gin"
)

func LoadLoginRouter(r *gin.RouterGroup) {
	r.POST("login", DoLogin)
}

// Token godoc
// @Summary      User authorization and generate token
// @Description  get a new token by login
// @Tags         Token
// @Accept       json
// @Produce      json
// @Param        Token body entities.User true "user name and password"
// @Success      201
// @Router       /login [post]
func DoLogin(ctx *gin.Context) {
	repoUser := &repository.User{}
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}

	dbuser, err := service.GetUserByName(repoUser, user.Name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
	}

	credentialError := utils.CheckPassword(dbuser.Password, user.Password)
	if credentialError != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		ctx.Abort()
		return
	}
	tokenString, err := utils.GenerateJWT(user.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
