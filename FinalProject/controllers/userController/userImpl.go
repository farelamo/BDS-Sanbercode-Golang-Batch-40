package userController

import (
	"fmt"
	"net/http"
	"FinalProject/models"
	. "FinalProject/services/userService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	userService UserService
}

func NewUserController(userService UserService) UserController {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) AddUser(ctx *gin.Context){
	var addUser models.AddUser

	if err := ctx.ShouldBindJSON(&addUser); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	
	user, err := u.userService.AddUser(&addUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (u *UserControllerImpl) GetUser(ctx *gin.Context) {
	users, err := u.userService.GetUser()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserControllerImpl) GetUserById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	user, err := u.userService.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err,
		})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserControllerImpl) UpdateUser(ctx *gin.Context) {
	var user models.User

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	count, err := u.userService.UpdateUser(id, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	message := fmt.Sprintf("Updated data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (u *UserControllerImpl) DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
		})
		return
	}

	count, err := u.userService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	message := fmt.Sprintf("Deleted data amount %d", count)
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})	
}


