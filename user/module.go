package user

import (
	"github.com/AntonBozhinov/techpodslearn/email"
	"github.com/gin-gonic/gin"
)

type Module struct {
	Repo        Repository
	EmailSender *email.Sender
	Router      *gin.Engine
}

func NewModule(repo Repository, emailSender *email.Sender, router *gin.Engine) *Module {
	return &Module{Repo: repo, EmailSender: emailSender, Router: router}
}

func (m *Module) Route() {
	m.Router.GET("/auth/google/login", m.Login)
	m.Router.GET("/auth/google/callback", m.LoginCallback)
}
