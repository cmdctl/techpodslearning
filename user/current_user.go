package user

import (
	"github.com/AntonBozhinov/techpodslearn/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (m *Module) getCurrentUser(c *gin.Context) {
	userValue, ok := c.Get("user")
	if !ok {
		errors.HTTP(c.Writer, "could not find user", http.StatusInternalServerError)
		return
	}
	user, ok := userValue.(*TokenClaims)
	if !ok {
		errors.HTTP(c.Writer, "could not extract user from session", http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}
