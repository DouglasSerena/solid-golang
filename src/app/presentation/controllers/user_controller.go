package controllers

import (
	"net/http"

	"github.com/DouglasSerena/solid-go/src/app/presentation/controllers/helpers"
	"github.com/DouglasSerena/solid-go/src/app/presentation/dto"
	"github.com/DouglasSerena/solid-go/src/app/usecase"
	"github.com/DouglasSerena/solid-go/src/domain"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	Find(context *gin.Context)
	FindBy(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type UserController struct {
	IUserController
	UserUsecase usecase.IUserUsecase
}

func (userController *UserController) Find(context *gin.Context) {
	users, err := userController.UserUsecase.FindAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.NewServerErrorRequest(err))
		return
	}
	context.JSON(http.StatusOK, helpers.NewOkRequest(users))
}

func (userController *UserController) FindBy(context *gin.Context) {
	user, err := userController.UserUsecase.FindBy(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.NewServerErrorRequest(err))
		return
	}
	context.JSON(http.StatusOK, helpers.NewOkRequest(user))
}

func (userController *UserController) Insert(context *gin.Context) {
	userDto := dto.UserCreateRequestDto{}
	err := context.ShouldBindJSON(&userDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.NewBadRequest(err))
		return
	}

	user, err := userController.UserUsecase.Create(&domain.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.NewServerErrorRequest(err))
	}
	context.JSON(http.StatusOK, helpers.NewOkRequest(dto.UserCreateResponseDto{Token: user.Token}))
}

func (userController *UserController) Update(context *gin.Context) {
	userDto := dto.UserCreateRequestDto{}
	err := context.ShouldBindJSON(&userDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, helpers.NewBadRequest(err))
		return
	}

	users, err := userController.UserUsecase.Update(&domain.User{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
	}, context.Param("id"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.NewServerErrorRequest(err))
		return
	}
	context.JSON(http.StatusOK, helpers.NewOkRequest(users))
}

func (userController *UserController) Delete(context *gin.Context) {
	err := userController.UserUsecase.Delete(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusInternalServerError, helpers.NewServerErrorRequest(err))
		return
	}
	context.JSON(http.StatusOK, helpers.NewOkRequest(nil))
}
