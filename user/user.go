package user

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const AllowedOrganization = "techpods.co"

type User struct {
	ID           bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Email        string        `json:"email" bson:"email"`
	ProfileImage string        `json:"profile_image" bson:"profile_image"`
	FirstName    string        `json:"first_name" bson:"first_name"`
	LastName     string        `json:"last_name" bson:"last_name"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
	HD            string `json:"hd"`
}

type TokenClaims struct {
	User
	jwt.StandardClaims
}
