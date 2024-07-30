package userDto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ListUserDto struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}
