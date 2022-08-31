package controllers

import (
	"net/http"
	"fmt"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hendricksakurvin/settings/api/responses"
	"github.com/hendricksakurvin/settings/api/models"

)

func (server *Server) Setting(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Here lies some settings and a bagel :)")
}

func (server *Server) GetSAllettings(w http.ResponseWriter, r *http.Request) {
	var setting = models.Setting{}
	fmt.Println(setting.GetSettings())
	responses.JSON(w, http.StatusOK, setting.GetSettings())
}

func (server *Server) GetSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err :=  strconv.ParseUint(vars["id"], 0, 64)

	if err != nil {
		fmt.Println(err)
		responses.JSON(w, http.StatusBadRequest, "im all out of love im so lost without you")
		return
	}
	
	var setting = models.Setting{}
	data := setting.GetSettingById(id)

	if data.Id == 0 {
		responses.JSON(w, http.StatusBadRequest, "notfound")
		return
	}

	responses.JSON(w, http.StatusOK, setting.GetSettingById(id))
}
