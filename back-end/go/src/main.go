package main

import (
	"daily_project_module/src/config/databases"
	"daily_project_module/src/routers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	
	app := gin.Default()
	routers.SetupRouters(app)
	dbErr := databases.CreateDBConnection()
	if dbErr != nil {
		fmt.Println(`Não foi possível conectar ao banco de dados: ${dbErr}`, dbErr)
	}
	app.Run()

}
