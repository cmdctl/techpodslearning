package user

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID           bson.ObjectId `json:"id"bson:"_id,omitempty"`
	Email        string        `json:"email"bson:"email"`
	ProfileImage string        `json:"profile_image"bson:"profile_image"`
	Token        string        `json:"token"bson:"token"`
	CreatedAt    time.Time     `json:"created_at"bson:"created_at"`
}
