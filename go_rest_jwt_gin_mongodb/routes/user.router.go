package routes

import (
	controller "github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb/controllers"
	middlewares "github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb/middlewares"

	"github.com/gin-gonic/gin"
)

/*make all the user routes protected via middleware.Authenticate()*/
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
