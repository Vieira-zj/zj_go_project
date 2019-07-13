package common

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	myutils "tools.app/utils"
)

// ******** Logger

// LogRequestData logs http resquest.
func LogRequestData(r *http.Request) error {
	reqData, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}
	log.Println("Request:\n", strings.Trim(string(reqData), "\n"))
	return nil
}

// LogDivLine logs division line.
func LogDivLine() {
	log.Println("|", strings.Repeat("*", 60))
}

// ******** Http Response

// JSONResponse json http response.
type JSONResponse struct {
	// Reserved field to add some meta information to the API response.
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

// JSONErrResponse json http error response.
type JSONErrResponse struct {
	Error *ErrorDesc `json:"error"`
}

// ErrorDesc json http error description.
type ErrorDesc struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
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
		Error: &ErrorDesc{Status: errCode, Desc: errMsg},
	})
}

// ErrHandler sends error response and logs error.
func ErrHandler(w http.ResponseWriter, err error) {
	WriteErrJSONResp(w, http.StatusInternalServerError, SvrErrRespMsg)
	log.Println(strings.Repeat("*", 6), err)
}

// ******** Parse Http Request Query

// GetStringArgFromQuery returns string value of args from request query form.
func GetStringArgFromQuery(r *http.Request, argName string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}

	if len(r.Form) > 0 {
		if val, ok := r.Form[argName]; ok {
			return val[0], nil
		}
	}
	return "", nil
}

// GetIntArgFromQuery returns int value of args from request query form.
func GetIntArgFromQuery(r *http.Request, argName string) (int, error) {
	val, err := GetStringArgFromQuery(r, argName)
	if err != nil {
		return -1, err
	}
	if len(val) == 0 {
		val = "-1"
	}

	ret, err := strconv.Atoi(val)
	if err != nil {
		return -1, err
	}
	return ret, err
}

// GetBoolArgFromQuery returns bool value of args from request query form.
func GetBoolArgFromQuery(r *http.Request, argName string) (bool, error) {
	val, err := GetStringArgFromQuery(r, argName)
	if err != nil {
		return false, err
	}
	if len(val) == 0 {
		val = "false"
	}

	ret, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return ret, nil
}

// ******** Helper functions

// CreateMockBytes returns mock md5 string for size of bytes.
func CreateMockBytes(size int) string {
	buf := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buf[i] = uint8(i % 64)
	}
	return myutils.GetBase64Text(buf)
}
