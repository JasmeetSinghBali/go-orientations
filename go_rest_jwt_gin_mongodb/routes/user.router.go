package routes

import(
	controller "go_rest_jwt_gin_mongodb/controllers"
	middlewares "go_rest_jwt_gin_mongodb/middlewares"
	"github.com/gin-gonic/gin"
)

/*make all the user routes protected via middleware*/
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}