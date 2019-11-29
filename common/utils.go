package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type appError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"status"`
}

type errorResource struct {
	Data appError `json:"data"`
}

// RespondWithJSON function called on sucess
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if response, err := json.Marshal(payload); err == nil {
		w.Write(response)
	}
}

// RespondWithError function called on error
func RespondWithError(w http.ResponseWriter, code int, message string, handlerError error) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HTTPStatus: code,
	}
	log.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if response, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(response)
	}
}
