package controller

import (
	"go-crud-api/model"
	"go-crud-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) userController {
	return userController{
		userUseCase: usecase,
	}
}

func (u *userController) GetUsers(ctx *gin.Context) {
	users, err := u.userUseCase.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *userController) CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)

	insertedUser, err := u.userUseCase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedUser)
}

func (u *userController) GetUsersByID(ctx *gin.Context) {
	userId := ctx.Param("userId")
	if userId == "" {
		response := model.Response{
			Message: "User ID can't is null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(userId)
	if err != nil {
		response := model.Response{
			Message: "User ID need int value",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := u.userUseCase.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	if user == nil {
		response := model.Response{
			Message: "User not finded in db",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
