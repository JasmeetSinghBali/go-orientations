package routes

import (
	controller "github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
