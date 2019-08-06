package learning

import "github.com/gin-contrib/cors"

func (s *Server) Cors() {
	s.Router.Use(cors.Default())
}
