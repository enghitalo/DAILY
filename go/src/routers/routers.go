package routers

import (
	AuthController "daily_project_module/src/modules/auth/controller"
	UserController "daily_project_module/src/modules/user/controller"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouters(app *gin.Engine) {
	app.GET("/", func(context *gin.Context) { context.JSON(200, gin.H{"hello": "world"}) })
	//jwtTestinho
	user := app.Group("/user")
	auth := app.Group("/auth")
	{
		user.POST("/create", UserController.CreateUser)
		user.GET("/getAll", UserController.GetAllUsers)
		user.GET("/:id", UserController.GetUserById)
	}
	{
		auth.POST("/login", AuthController.Login)
	}
}
