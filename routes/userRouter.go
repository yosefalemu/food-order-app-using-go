package routes

import (
	"golang-food-store-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine) {
	// Use a reference to the GetUsers function, don't call it directly
	incomingRoutes.POST("/users/register", controllers.RegisterUser)
	incomingRoutes.POST("/users/login", controllers.LoginUser)
	incomingRoutes.GET("/users", controllers.GetUsers)
	incomingRoutes.GET("/users/:user_id", controllers.GetSingleUser)
	incomingRoutes.GET("/users/getuserprofile/:user_id", controllers.GetUserProfile)
	incomingRoutes.PATCH("/users/updateuser/:user_id", controllers.UpdateUser)
	incomingRoutes.DELETE("/users/:user_id", controllers.DeleteUser)

}
