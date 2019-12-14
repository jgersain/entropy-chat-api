package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

//Return json responses
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

//Return error json responses
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
	return
}

func FormatError(err string) error {

	if strings.Contains(err, "email") {
		return errors.New("el correo electrónico ya ha sido registrado")
	}

	return errors.New("valores no válidos")
}
