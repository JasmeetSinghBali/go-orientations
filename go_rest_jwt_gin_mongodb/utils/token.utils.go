package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go_rest_jwt_gin_mongodb/database"
	jwt "github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Signature payload for jwt*/
type signPD struct{
	Email 		string
	First_name 	string
	Last_name 	string
	Uid 		string
	User_Type   string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

/*Generate JWT Token*/
func GenerateJWTTokens(email string, firstName string, lastName string, userType string, uid string)(signedToken string signedRefreshToken string, err error){
	claims = &signPD{
		Email: email,
		First_name: firstName,
		Last_name: lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &signPD{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	refreshToken ,err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SECRET_KEY))

	if err != nil{
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}
