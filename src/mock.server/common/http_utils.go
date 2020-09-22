package common

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"regexp"
	"strconv"
	"strings"
)

/* Logger */

// LogRequestData logs http resquest.
func LogRequestData(r *http.Request) error {
	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		return err
	}

	log.Println("Request:\n", strings.Trim(string(req), "\n"))
	return nil
}

// LogDivLine logs division line.
func LogDivLine() {
	log.Println("|", strings.Repeat("*", 60))
}

/* Http Response */

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

// WriteOKJSONResp writes http ok response as a standard JSON.
func WriteOKJSONResp(w http.ResponseWriter, data interface{}) error {
	w.Header().Set(TextContentType, ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(&JSONResponse{Data: data})
}

// WriteOKHTMLResp writes http ok response as html.
func WriteOKHTMLResp(w http.ResponseWriter, data []byte) error {
	w.Header().Set(TextContentType, ContentTypeHTML)
	w.Header().Set(TextContentLength, strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		return err
	}
	return nil
}

// AddCorsHeaders writes headers to fix CORS issue.
func AddCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept,Origin,Content-Type,X-Custom-Header")
}

// WriteErrJSONResp writes http error response as a Standard API JSON with a resp code.
func WriteErrJSONResp(w http.ResponseWriter, errCode int, errMsg string) error {
	w.Header().Set(TextContentType, ContentTypeJSON)
	w.WriteHeader(errCode)
	return json.NewEncoder(w).Encode(
		&JSONErrResponse{Error: &ErrorDesc{Status: errCode, Desc: errMsg}})
}

// ErrHandler handles "internal server error".
func ErrHandler(w http.ResponseWriter, err error) {
	log.Println(strings.Repeat("*", 6), err)
	if err := WriteErrJSONResp(w, http.StatusInternalServerError, SvrErrRespMsg); err != nil {
		log.Println(strings.Repeat("*", 6), err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

/* Parse Http Request Query */

// GetStringArgFromQuery returns string value of arg from request query form.
func GetStringArgFromQuery(r *http.Request, argName string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}

	if len(r.Form) > 0 {
		if val, ok := r.Form[argName]; ok {
			return val[0], nil
		}
	}
	return "", fmt.Errorf("key [%s] not exist", argName)
}

// GetIntArgFromQuery returns int value of arg from request query form.
func GetIntArgFromQuery(r *http.Request, argName string) (int, error) {
	val, err := GetStringArgFromQuery(r, argName)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(val)
}

// GetBoolArgFromQuery returns bool value of args from request query form.
func GetBoolArgFromQuery(r *http.Request, argName string) (bool, error) {
	val, err := GetStringArgFromQuery(r, argName)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(val)
}

/* Mock API, Template Functions */

// ParseParamsForTempl parse query params for templated response.
func ParseParamsForTempl(query map[string][]string) (map[string]string, error) {
	values := make(map[string]string, len(query))
	for k, v := range query {
		val := v[0]
		if strings.Contains(val, "randint") {
			num, err := getNumberArg(val)
			if err != nil {
				return nil, err
			}
			val = strconv.Itoa(rand.Intn(num))
		} else if strings.Contains(val, "randstr") {
			num, err := getNumberArg(val)
			if err != nil {
				return nil, err
			}
			val = CreateMockString(num)
		} else if strings.Contains(val, "randchoice") {
			strArgs := val[strings.Index(val, "(")+1 : len(val)-1]
			args := strings.Split(strArgs, ",")
			val = args[rand.Intn(len(args))]
		}

		values[k] = val
	}

	return values, nil
}

func getNumberArg(text string) (int, error) {
	r, err := regexp.Compile(`\d+`)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(r.FindString(text))
}

// QueryToMap formats string query to map[string][]string (compatible with r.URL.Query()).
func QueryToMap(query string) map[string][]string {
	items := strings.Split(query, "&")
	retMap := make(map[string][]string, len(items))

	for _, item := range items {
		tmp := strings.Split(item, "=")
		retMap[tmp[0]] = []string{tmp[1]}
	}
	return retMap
}
