package main

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, r *http.Request, v interface{}) {
	jsonByte, errJSON := json.Marshal(v)
	if errJSON != nil {

	}

	// json header
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonByte)
}

func writeErrorWithMessage(w http.ResponseWriter, r *http.Request, httpCode int, msg string) {
	jsonByte, errJSON := json.Marshal(
		struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    httpCode,
			Message: msg,
		},
	)
	if errJSON != nil {

	}

	// json header
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpCode)
	w.Write(jsonByte)
}

func writeError(w http.ResponseWriter, r *http.Request, httpCode int) {
	w.WriteHeader(httpCode)
}
