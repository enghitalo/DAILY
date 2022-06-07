package controller

import (
	authService "daily_project_module/src/modules/auth/service"
	"daily_project_module/src/modules/user/dto"
	userService "daily_project_module/src/modules/user/service"

	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(context *gin.Context) {
	createUserDto := dto.CreateUserDto{}

	if err := context.BindJSON(&createUserDto); err != nil {
		context.AbortWithError(http.StatusUnprocessableEntity, err)
		return
	}

	bcrypt_hash, err := bcrypt.GenerateFromPassword([]byte(createUserDto.Password), bcrypt.MinCost)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	createUserDto.Password = string(bcrypt_hash)
	user, err := userService.CreateUser(createUserDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusCreated, user)
}

func GetAllUsers(context *gin.Context) {
	token := context.GetHeader("token")

	if err := authService.AuthVerify(token); err != nil {
		context.Status(http.StatusUnauthorized)
		return
	}
	users, err := userService.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserById(context *gin.Context) {
	params := context.Params

	id, parse_err := strconv.Atoi(params.ByName("id"))

	if parse_err == nil {
		fmt.Println("Erro:", parse_err)
	}

	user, err := userService.GetUserById(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	context.JSON(http.StatusOK, user)
}
