package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `bson:"name" json:"name" binding:"required" error:"Field name is required"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password" binding:"required"`
	IsDeleted bool               `bson:"isDeleted" json:"isDeleted"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}
