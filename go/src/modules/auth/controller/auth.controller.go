package controller

import (
	"daily_project_module/src/modules/user/service"
	util "daily_project_module/src/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:",omitempty" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	UserCredential
}

func Login(context *gin.Context) {
	userCredential := UserCredential{}

	if err := context.BindJSON(&userCredential); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// userFilter control
	user, err := service.GetUserToAuthentication(userCredential)
	if err != nil {
		context.AbortWithError(http.StatusBadRequest, fmt.Errorf("usuário não encontrado %v", err))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredential.Password)); err != nil {
		context.AbortWithError(http.StatusBadRequest, fmt.Errorf("a senha não confere ao login do usuário %v", err))
		return
	}

	userCredential.Password = ""

	/// convert Struct fields into Map String
	var myMap map[string]interface{}
	data, _ := json.Marshal(userCredential)
	json.Unmarshal(data, &myMap)
	token, _ := util.EncodeJwtToken(myMap)

	context.JSON(http.StatusOK, gin.H{"token": token, "user": userCredential})
}
