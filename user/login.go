package user

import (
	"github.com/AntonBozhinov/techpodslearn/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func (m *Module) Login(c *gin.Context) {
	OAuthGoogleLogin(c.Writer, c.Request)
}

func (m *Module) LoginCallback(c *gin.Context) {
	googleUser, err := OAuthGoogleCallback(c.Writer, c.Request)
	if err != nil {
		errors.HTTP(c.Writer, "login failed", http.StatusInternalServerError)
		return
	}
	if googleUser.HD != AllowedOrganization {
		errors.HTTP(c.Writer, "Please, login with your Techpods email", http.StatusBadRequest)
	}
	user := &User{
		Email:        googleUser.Email,
		ProfileImage: googleUser.Picture,
		FirstName:    googleUser.GivenName,
		LastName:     googleUser.FamilyName,
		CreatedAt:    time.Now(),
	}

	exist, err := m.Repo.Exist(user.Email)
	if err != nil {
		errors.HTTP(c.Writer, "could not check user", http.StatusInternalServerError)
		return
	}
	if exist {
		read, err := m.Repo.Read(user.Email)
		if err != nil {
			errors.HTTP(c.Writer, "could not get user", http.StatusInternalServerError)
			return
		}
		token, err := createJWT(read)
		if err != nil {
			errors.HTTP(c.Writer, "could not create session jwt", http.StatusInternalServerError)
			return
		}
		http.Redirect(c.Writer, c.Request, "/?token=" + token, http.StatusTemporaryRedirect)
	} else {
		_, err := m.Repo.Create(user)
		if err != nil {
			errors.HTTP(c.Writer, "could not create user", http.StatusInternalServerError)
			return
		}
		token, err := createJWT(user)
		if err != nil {
			errors.HTTP(c.Writer, "could not create session jwt", http.StatusInternalServerError)
			return
		}
		http.Redirect(c.Writer, c.Request, "/?token=" + token, http.StatusTemporaryRedirect)
	}
}

func sendGreetingEmail() {

}

func createJWT(user *User) (string, error) {
	claims := TokenClaims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 1, 0).Unix(),
			Issuer:    AllowedOrganization,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("could not create session jwt error: " + err.Error())
	}
	return str, nil
}
