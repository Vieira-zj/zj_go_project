package utils

import (
	"encoding/json"
	"net/http"
)

// ********* HTTP Response

// JSONResponse json http response
type JSONResponse struct {
	// Reserved field to add some meta information to the API response.
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// JSONErrResponse json http error response
type JSONErrResponse struct {
	Error *APIError `json:"error"`
}

// APIError json http error response
type APIError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

// WriteOKJSONResp writes the response as a standard JSON response with StatusOK.
func WriteOKJSONResp(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: m}); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteOKHTMLResp returns html with StatusOK.
func WriteOKHTMLResp(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteErrJSONResp writes the error response as a Standard API JSON response with a response code.
func WriteErrJSONResp(w http.ResponseWriter, errCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&JSONErrResponse{Error: &APIError{Status: errCode, Title: errMsg}})
}
