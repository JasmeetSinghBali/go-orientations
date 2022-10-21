package middleware

import (
	"fmt"
	"net/http"

	utils "github.com/Jasmeet-1998/go-orientations/go_rest_jwt_gin_mongodb/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No token was supplied in auth header.")})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_Type)
		c.Next()
	}
}
