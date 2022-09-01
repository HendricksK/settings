package controllers

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/hendricksakurvin/settings/api/responses"
	"github.com/hendricksakurvin/settings/api/models"

)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Baby come back, any kind of fool could see " +
		"There was something in everything about you " +
		"Baby come back, you can blame it all on me " +
		"I was wrong and I just can't live without you")
}

func (server *Server) GetAllSettings(w http.ResponseWriter, r *http.Request) {
	var setting = models.Setting{}
	fmt.Println(setting.GetAll())
	responses.JSON(w, http.StatusOK, setting.GetAll())
}

func (server *Server) ReadSetting(w http.ResponseWriter, r *http.Request) {
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

	responses.JSON(w, http.StatusOK, data)
}

func (server *Server) CreateSetting(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	setting := models.Setting{}
	err = json.Unmarshal(body, &setting)

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, "im all out of love im so lost without you")
	}

	settingCreated, err := setting.Create(server.DB)

	responses.JSON(w, http.StatusOK, settingCreated)
}

func (server *Server) UpdateSetting(w http.ResponseWriter, r *http.Request) {
	var setting = models.Setting{}
	responses.JSON(w, http.StatusOK, setting.Update())
}

func (server *Server) DeleteSetting(w http.ResponseWriter, r *http.Request) {
	var setting = models.Setting{}
	responses.JSON(w, http.StatusOK, setting.Delete())
}
