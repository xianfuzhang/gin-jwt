package adapter

import (
	"net/http"
	"test/v2/internal/adapter/repository"
	"test/v2/internal/application"
	"test/v2/internal/entities"
	"test/v2/internal/types"
	"test/v2/internal/utils"

	"github.com/gin-gonic/gin"
)

var repoUser = &repository.User{}

func LoadUserRouter(r *gin.RouterGroup) {
	r.POST("users", CreateUser)
	r.PUT("users/:userName/reset", ResetPassword)
}

// CreateUser godoc
// @Summary      Create user account
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        User body entities.User true "user name and password"
// @Success      201
// @Router       /v1/users [post]
func CreateUser(ctx *gin.Context) {
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

// @Summary      Reset user password
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param 		 user_name path string true "User Name"
// @Success      204
// @Router       /v1/users/{user_name}/reset [put]
func ResetPassword(ctx *gin.Context) {
	var (
		user    entities.User
		hashPwd string
		err     error
	)
	if user, err = application.GetUserByName(repoUser, ctx.Param("userName")); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if hashPwd, err = hashPassword(types.UserResetPwd); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user.Password = hashPwd
	err = application.UpdateUserPassword(repoUser, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func hashPassword(pwd string) (string, error) {
	return utils.HashPassword(pwd)
}
