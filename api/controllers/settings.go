package controllers

import (
	"net/http"

	"github.com/hendricksakurvin/settings/api/responses"
	"github.com/hendricksakurvin/settings/api/entities/setting"
)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Here lies some settings and a bagel :)")
	// responses.JSON(w, http.StatusOK, )
}

// func getSettingById() {
// 	var data = setting.Get
// }