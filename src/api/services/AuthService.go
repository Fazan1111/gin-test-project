package services

import (
	"context"
	"fmt"
	"learnGin/src/api/models"
	"learnGin/src/loader/mongo"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RunAuth(c *gin.Context, id string) models.User {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Error converting ObjectId:", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Error converting ObjectId"})
	}

	var result bson.M

	collErr := mongo.DB.Collection(UserCollName).FindOne(context.TODO(), bson.D{{Key: "_id", Value: objID}}).Decode(&result)

	fmt.Println("run auth error", err)
	if collErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
	}

	fmt.Println("result auth", result)
	resp := models.User{
		Id:    result["_id"].(primitive.ObjectID),
		Name:  result["name"].(string),
		Email: result["email"].(string),
	}
	fmt.Println("resp auth", resp)
	return resp
}

func FindUserByMail(c *gin.Context, mail string) models.User {
	var user bson.M
	err := mongo.DB.Collection(UserCollName).FindOne(context.TODO(), bson.D{{Key: "email", Value: mail}}).Decode(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not found"})
	}

	resp := models.User{
		Id:       user["_id"].(primitive.ObjectID),
		Name:     user["name"].(string),
		Email:    user["email"].(string),
		Password: user["password"].(string),
	}

	return resp
}
