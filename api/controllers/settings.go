package controllers

import (
	"net/http"

	"github.com/hendricksakurvin/settings/api/responses"
)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}