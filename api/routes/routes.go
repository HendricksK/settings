package routes

import (
	"fmt"

	"github.com/hendricksakurvin/settings/api/middleware"
)

func (s *Server) initializeRoutes() {
	fmt.Println("bagel routes")
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")
}