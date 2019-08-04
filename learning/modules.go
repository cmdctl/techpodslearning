package learning

import "github.com/AntonBozhinov/techpodslearn/user"

func (s *Server) Modules() {
	s.UserModule()
}

func (s *Server) UserModule() {
	repository := user.NewMongoRepository(s.Mongo)
	module := user.NewModule(repository, s.EmailSender, s.Router)
	module.Route()
}
