package controllers

import (
	"net/http"

	"github.com/hendricksakurvin/settings/api/responses"
	"github.com/hendricksakurvin/settings/api/entities"
)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Here lies some settings and a bagel :)")
}