package userModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User Schema [DTO] for conversion json<>golang primitive DataTypes*/
//📝 better to use bson(binary JSON) as it is binary representation of JSON data and for languages like golang which have complex primitive data types bson helps in performing precise calculations on the data
// NOTE- internally and over the network mongodb actually stores data in BSON only which offer better parsing and querying,indexing performance.
type User struct {
	ID primitive.ObjectID `bson:"_id"`
	// 📝this will be stored in db as first_name while getting data from client it will be as key First_name
	First_Name *string `json: "first_name" validate: "required, min=6, max=32"`
	Last_Name  *string `json: "last_name" validate: "required, min=6,max=32"`
	Password   *string `json: "password" validate: "required, min=9,max=60"`
	Email      *string `json: "email" validate: "email, required"`
	Phone      *string `json: "phone_no" validate: "required"`
	//📝 just like enum in javascript where u can define the possible types a schema key can posses we have 'eq=' aka equal to in golang
	User_Type     *string   `json: "user_type" validate: "required, eq=ADMIN|eq=USER"`
	Token         *string   `json: "token"`
	Refresh_Token *string   `json: "refresh_token"`
	Created_At    time.Time `json: "created_at"`
	Updated_At    time.Time `json: "updated_at"`
	User_ID       *string   `json: "user_id"`
}
