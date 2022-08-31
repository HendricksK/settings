package controllers

import (
	"net/http"
	"fmt"

	"github.com/hendricksakurvin/settings/api/responses"
	"github.com/hendricksakurvin/settings/api/models"
)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Here lies some settings and a bagel :)")
}

func (server *Server) GetSetting(w http.ResponseWriter, r *http.Request) {
	var setting = models.Setting{}
	fmt.Println(setting.GetSettingById(12))
	responses.JSON(w, http.StatusOK, setting.GetSettingById(12))
}
