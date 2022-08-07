package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func InspectUserTypeToUid(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	// check role ADMIN or not
	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	// returns nil as error
	return err
}

func CheckUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	/*User is not admin and cannot access data of uid that belong to other user*/
	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	/*check if user is admin and want to have specific userid details then allow*/
	err = InspectUserTypeToUid(c, userType)
	return err
}
