package user

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID           bson.ObjectId `json:"id"bson:"_id,omitempty"`
	Email        string        `json:"email"`
	ProfileImage string        `json:"profile_image"`
	Token        string        `json:"token"`
	CreatedAt    time.Time     `json:"created_at"`
}
