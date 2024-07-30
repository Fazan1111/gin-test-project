package authDto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegistorResp struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	IsDeleted    bool               `json:"isDeleted"`
	CreatedAt    time.Time          `json:"createdAt"`
	AccessToken  string             `json:"accessToken"`
	RefreshToken string             `json:"refreshToken"`
}
