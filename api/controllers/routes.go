package controllers

import (
	"fmt"

	"github.com/hendricksakurvin/settings/api/middleware"
)

func (s *Server) initializeRoutes() {
	fmt.Println("bagel routes")
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Setting)).Methods("GET")
	s.Router.HandleFunc("/yes", middleware.SetMiddlewareJSON(s.GetSetting)).Methods("GET")
	// s.Router.HandleFunc("/setting", middleware.SetMiddlewareJSON(s.GetSetting)).Methods("GET")
}