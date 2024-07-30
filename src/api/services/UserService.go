package services

import (
	"context"
	"fmt"
	userDto "learnGin/src/api/dto/user"
	"learnGin/src/api/models"
	customerError "learnGin/src/common/customError"
	"learnGin/src/loader/mongo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoLib "go.mongodb.org/mongo-driver/mongo"
)

var UserCollName string = "users"

type UserService struct {
	// Any dependencies your service might have can be added here
}

func FindUserByName(c *gin.Context, name string) any {
	var ctx = context.Background()
	var user models.User
	err := mongo.DB.Collection(UserCollName).FindOne(ctx, bson.D{{Key: "name", Value: name}}).Decode(&user)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return user
}

func ListUsers() any {
	ctx := context.TODO()
	pipeline := mongoLib.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{Key: "isDeleted", Value: false},
				},
			},
		},
		bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "_id", Value: 1},
					{Key: "name", Value: 1},
					{Key: "email", Value: 1},
					{Key: "createdAt", Value: 1},
				},
			},
		},
	}
	coll, err := mongo.DB.Collection(UserCollName).Aggregate(ctx, pipeline)
	if err != nil {
		fmt.Println("Cannot fetch users")
		return nil
	}

	var result []userDto.ListUserDto
	if err = coll.All(ctx, &result); err != nil {
		return nil
	}

	return result
}

func CreateUser(c *gin.Context, user models.User) models.User {
	var ctx = context.Background()
	newUser, err := mongo.DB.Collection(UserCollName).InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			customerError.ResponseError(customerError.CANNT_CREATE_DOC, "Can't create user"),
		)
	}

	user.Id = newUser.InsertedID.(primitive.ObjectID)
	return user
}
