package common

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/golib/httprouter"
)

// ******** Log

// PerfLogger print log with handler run time.
func PerfLogger(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		LogDivLine()
		log.Printf("Start: %s %s\n", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("Done (%s %s): %v\n", r.Method, r.URL.Path, time.Since(start))
		LogDivLine()
	}
}

// LogRequestData logs http resquest.
func LogRequestData(r *http.Request) error {
	reqData, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	log.Println("Request:\n", string(reqData))
	return nil
}

// LogDivLine logs division line.
func LogDivLine() {
	log.Println("|", strings.Repeat("*", 60))
}

// ******** Http response

// JSONResponse json http response
type JSONResponse struct {
	// Reserved field to add some meta information to the API response
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
	w.Header().Set(TextContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: m}); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, SvrErrRespMsg)
	}
}

// WriteOKHTMLResp returns html with StatusOK.
func WriteOKHTMLResp(w http.ResponseWriter, data []byte) {
	w.Header().Set(TextContentType, ContentTypeHTML)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, SvrErrRespMsg)
	}
}

// WriteErrJSONResp writes the error response as a Standard API JSON response with a response code.
func WriteErrJSONResp(w http.ResponseWriter, errCode int, errMsg string) {
	w.Header().Set(TextContentType, ContentTypeJSON)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&JSONErrResponse{
		Error: &APIError{Status: errCode, Title: errMsg},
	})
}

// ErrHandler sends error response and logs error.
func ErrHandler(w http.ResponseWriter, err error) {
	WriteErrJSONResp(w, http.StatusInternalServerError, SvrErrRespMsg)
	log.Println(strings.Repeat("*", 6), err)
}
