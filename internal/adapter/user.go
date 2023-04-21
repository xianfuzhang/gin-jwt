package adapter

import (
	"net/http"
	"test/v2/internal/adapter/repository"
	"test/v2/internal/application"
	"test/v2/internal/entities"
	"test/v2/internal/utils"

	"github.com/gin-gonic/gin"
)

func LoadUserRouter(r *gin.RouterGroup) {
	r.POST("users", CreateUser)
}

// func RegisterUser(context *gin.Context) {
// 	var user entities.User
// 	if err := context.ShouldBindJSON(&user); err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}
// 	if _, err := utils.HashPassword(user.Password); err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		context.Abort()
// 		return
// 	}
// 	record := database.Instance.Create(&user)
// 	if record.Error != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
// 		context.Abort()
// 		return
// 	}
// 	context.JSON(http.StatusCreated, gin.H{"username": user.Name})
// }

// CreateUser godoc
// @Summary      Create user account
// @Description  create user account
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Param        User body entities.User true "user name and password"
// @Success      201
// @Router       /users [post]
func CreateUser(ctx *gin.Context) {
	repoUser := &repository.User{}
	var (
		user    entities.User
		hashPwd string
		err     error
	)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if hashPwd, err = utils.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user.Password = hashPwd
	err = application.CreateUser(repoUser, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"username": user.Name})
}
