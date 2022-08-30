package controllers

import (
	"net/http"

	"github.com/hendricksakurvin/settings/api/responses"
)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Here lies some settings and a bagel :)")

}