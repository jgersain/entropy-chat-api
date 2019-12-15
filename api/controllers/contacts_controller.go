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

func (server *Server) CreateContact(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	contact := models.Contact{}
	err = json.Unmarshal(body, &contact)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = contact.Validate()
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}
	if uid != contact.UserID {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}
	contactCreated, err := contact.SaveContact(server.DB)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		utils.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, contactCreated.ID))
	utils.JSON(w, http.StatusCreated, contactCreated)
}

func (server *Server) GetContactsUser(w http.ResponseWriter, r *http.Request) {

	uid, err := strconv.ParseUint(r.FormValue("user_id"), 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	contact := models.Contact{}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}
	if tokenID != uint32(uid) {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}

	contacts, err := contact.FindAllContactsUser(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, contacts)
}

func (server *Server) GetContactUser(w http.ResponseWriter, r *http.Request) {

	cId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	uId, err := strconv.ParseUint(r.FormValue("user_id"), 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	contact := models.Contact{}

	tokenID, err := auth.ExtractTokenID(r)
	if err != nil {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}
	if tokenID != uint32(uId) {
		utils.ERROR(w, http.StatusUnauthorized, errors.New("No autorizado"))
		return
	}

	contactReceived, err := contact.FindContactUserByID(server.DB, uint32(uId), uint32(cId))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.JSON(w, http.StatusOK, contactReceived)
}
