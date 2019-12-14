package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jgersain/entropy-chat-api/api/auth"
	"github.com/jgersain/entropy-chat-api/api/models"
	"github.com/jgersain/entropy-chat-api/api/utils"
	"io/ioutil"
	"net/http"
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
