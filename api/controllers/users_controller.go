package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jgersain/entropy-chat-api/api/auth"
	"github.com/jgersain/entropy-chat-api/api/models"
	"github.com/jgersain/entropy-chat-api/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	//read body
	body, err := ioutil.ReadAll(r.Body)
	user := models.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("")
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	utils.JSON(w, http.StatusCreated, userCreated)
}

func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}
	if tokenID != uint32(uid) {
		utils.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	err = user.Validate("")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, updatedUser)
}
