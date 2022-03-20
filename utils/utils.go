package utils

import (
	"log"
	"net/http"
)

// GenericFatalErrorHandler is a generic way to kill the process if there was a fatal error.
func GenericFatalErrorHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// GenericAPIErrorHandler returns a 400 API error with a custom statement if required.
func GenericAPIErrorHandler(err error, w http.ResponseWriter, error_message string) bool {
	if err != nil {
		w.WriteHeader(400)
		_, writeError := w.Write([]byte(error_message + " " + err.Error()))
		if writeError != nil {
			log.Fatal(writeError.Error())
		}
		return true
	}
	return false
}

// GenericAPI400Response returns a 400 API error with a custom statement.
func GenericAPI400Response(w http.ResponseWriter, error_message string) {
	w.WriteHeader(400)
	_, writeError := w.Write([]byte(error_message + " "))
	if writeError != nil {
		log.Fatal(writeError.Error())
	}
}
