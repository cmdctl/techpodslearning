package learning

import (
	"github.com/AntonBozhinov/techpodslearn/email"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"net/http"
)

type Server struct {
	Router      *gin.Engine
	Mongo       *mgo.Database
	EmailSender *email.Sender
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func NewServer(router *gin.Engine, mongo *mgo.Database, emailSender *email.Sender) *Server {
	return &Server{Router: router, Mongo: mongo, EmailSender: emailSender}
}

