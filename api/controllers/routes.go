package controllers

import (
	"fmt"

	"github.com/hendricksakurvin/settings/api/middleware"
)

func (s *Server) initializeRoutes() {
	fmt.Println("bagel routes")
	s.Router.HandleFunc("/", middleware.SetMiddlewareJSON(s.Home)).Methods("GET")
	s.Router.HandleFunc("/settings", middleware.SetMiddlewareJSON(s.GetAllSettings)).Methods("GET")
	s.Router.HandleFunc("/settings/{id}", middleware.SetMiddlewareJSON(s.ReadSetting)).Methods("GET")
	s.Router.HandleFunc("/settings", middleware.SetMiddlewareJSON(s.CreateSetting)).Methods("POST")
	s.Router.HandleFunc("/settings", middleware.SetMiddlewareJSON(s.UpdateSetting)).Methods("PUT")
	s.Router.HandleFunc("/settings", middleware.SetMiddlewareJSON(s.DeleteSetting)).Methods("DELETE")
}