package service

import (
	"daily_project_module/src/config/databases"
	"daily_project_module/src/modules/user/dto"
	"daily_project_module/src/modules/user/models"
	"encoding/json"
	"fmt"
)

func CreateUser(dto dto.CreateUserDto) (any, error) {
	dbConn, dbErr := databases.GetDatabaseConnection()
	if dbErr != nil {
		fmt.Println(dbConn.Error)
		return nil, dbErr
	}

	var user *models.User = &models.User{
		Username: dto.Username,
		Password: dto.Password,
	}

	db_instance := dbConn.Create(user)

	return user, db_instance.Error

}

func GetAllUsers() (any, error) {
	dbConn, dbErr := databases.GetDatabaseConnection()
	if dbErr != nil {
		fmt.Println(dbConn.Error)
		return nil, dbErr
	}

	users := []models.User{}

	db_instance := dbConn.Find(&users)

	return users, db_instance.Error

}

func GetUserById(id int) (any, error) {
	dbConn, dbErr := databases.GetDatabaseConnection()
	if dbErr != nil {
		fmt.Println(dbConn.Error)
		return nil, dbErr
	}

	user := &models.User{}

	// user

	db_instance := dbConn.First(user, id)

	return user, db_instance.Error

}

func GetUserToAuthentication(userFilter any) (*models.User, error) {
	var myMap map[string]interface{}
	data, _ := json.Marshal(userFilter)
	json.Unmarshal(data, &myMap)

	dbConn, dbErr := databases.GetDatabaseConnection()
	if dbErr != nil {
		return nil, dbErr
	}

	user := &models.User{}

	db_instance := dbConn.First(&user, map[string]interface{}{
		"active":   true,
		"username": myMap["username"],
	})

	return user, db_instance.Error
}
