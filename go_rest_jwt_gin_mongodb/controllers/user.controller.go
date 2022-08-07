package controllers

import(
	"context"
	"fmt"
	"log"
	"strconv"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	utils "go_rest_jwt_gin_mongodb/utils"
	"go_rest_jwt_gin_mongodb/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validate.new()

func HashPassword() {

}

func VerifyPassword() {

}

func Signup() *gin.HandlerFunc {
	return func(c *gin.Context){
		var ctx, cancel := context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		/*grab and store the new user as JSON in user var from request via c gin context*/
		if err := c.BindJSON(&user); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// validate the request payload data
		validationErr := validate.Struct(user)
		if validationErr != nil{
			c.JSON(http.StatusBadRequest,gin.H{"error":validationErr.Error()})
			return
		}

		// check for duplicate email
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occured while inspecting user email exists or not"})
		}

		//check for duplicate phone
		count, err := userCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
		defer cancel()
		if err != nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"phone":"error occured while while inspecting user phone exists or not"})
		}

		if count > 0{
			c.JSON(http.StatusInternalServerError,gin.H{"error": "this email or user already exists"})
		}
	}
}

func Login() {

}

func GetUsers(){

}

/*Get user by userID only for admin if userid differs from the current userId*/
func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := utils.CheckUserTypeToUid(c, userId); err !=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var ctx, cancel := context.WithTimeout(context.Background(),100*time.Second)
		var user models.User
		/*find & grab the user id data in user var*/
		// Decode to convert the json into primitive data structure so golang understands it
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			return
		}
		C.JSON(http.StatusOK,user)	
	}	
}
