package adapter

import (
	"net/http"
	"strconv"
	"test/v2/internal/adapter/repository"
	"test/v2/internal/entities"
	"test/v2/internal/service"
	"test/v2/internal/types"
	"test/v2/internal/utils"

	"github.com/gin-gonic/gin"
)

var repoUser = &repository.User{}

func LoadUserRouter(r *gin.RouterGroup) {
	r.POST("users", createUser)
	r.GET("users", getUsers)
	r.DELETE("users/:userId", deleteUser)
	r.PUT("users/:userName/reset", resetPassword)
}

// @Summary      Get user list
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        limit  query integer true "Limit" default(10)
// @Param        offset query integer true "Offset" default(0)
// @Success      200 {array} entities.User
// @Router       /v1/users [get]
func getUsers(ctx *gin.Context) {
	var (
		users  []entities.User
		err    error
		limit  int64
		offset int64
	)
	limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	offset, err = strconv.ParseInt(ctx.Query("offset"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	users, err = service.FetchUsers(repoUser, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// @Summary      Create user account
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param        User body entities.User true "user name and password"
// @Success      201
// @Router       /v1/users [post]
func createUser(ctx *gin.Context) {
	var (
		user    entities.User
		exist   entities.User
		hashPwd string
		err     error
	)
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if exist, _ = service.GetUserByName(repoUser, user.Name); exist.Name != "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "user name exists"})
		ctx.Abort()
		return
	}
	if hashPwd, err = hashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	user.Password = hashPwd
	err = service.CreateUser(repoUser, &user)
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
func resetPassword(ctx *gin.Context) {
	var (
		user    entities.User
		hashPwd string
		err     error
	)
	if user, err = service.GetUserByName(repoUser, ctx.Param("userName")); err != nil {
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
	if err = service.UpdateUserPassword(repoUser, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// @Summary      Delete user
// @Tags         User Account
// @Accept       json
// @Produce      json
// @Security     Bearer
// @Param 		 user_id path integer true "User ID"
// @Success      204
// @Router       /v1/users/{user_id} [delete]
func deleteUser(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	if err := service.DeleteUser(repoUser, userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func hashPassword(pwd string) (string, error) {
	return utils.HashPassword(pwd)
}
