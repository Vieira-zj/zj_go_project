package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	// IsHTTPLog is print logs for http request and response.
	IsHTTPLog = true
)

// ********* HTTP Request (client)

// HTTPRequest model for http request
type HTTPRequest struct {
	URL     string
	Method  string
	Headers http.Header
	Query   url.Values
	Body    string
}

// HTTPGet sends http GET request, and returns response.
func HTTPGet(params *HTTPRequest) (*http.Response, error) {
	params.Method = "GET"
	printHTTPRequest(params)

	u, err := url.Parse(params.URL)
	if err != nil {
		printErrorLine(err)
		return nil, err
	}
	if len(params.Query) > 0 {
		u.RawQuery = params.Query.Encode()
	}

	resp, err := http.Get(u.String())
	if err != nil {
		printErrorLine(err)
		return nil, err
	}
	defer resp.Body.Close()
	if err := printHTTPResponse(params.URL, resp); err != nil {
		printErrorLine(err)
		return nil, err
	}

	return resp, nil
}

// HTTPPost sends http Post request, and returns response.
func HTTPPost(params *HTTPRequest) (*http.Response, error) {
	params.Method = "POST"
	printHTTPRequest(params)

	req, err := http.NewRequest(params.Method, params.URL, strings.NewReader(params.Body))
	if err != nil {
		printErrorLine(err)
		return nil, err
	}
	req.Header = params.Headers

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		printErrorLine(err)
		return nil, err
	}
	defer resp.Body.Close()

	if err := printHTTPResponse(params.URL, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func printHTTPRequest(params *HTTPRequest) {
	url := params.URL
	if len(params.Query) > 0 {
		url = url + "?" + params.Query.Encode()
	}
	printDivLine()
	printlnPrefixLine(fmt.Sprintf("-Request: (%s) %s", params.Method, url))
	printDivLine()

	if len(params.Headers) > 0 {
		printlnPrefixLine("-Headers:")
		for k, v := range params.Headers {
			printlnPrefixLine(fmt.Sprintf("%s: %s", k, strings.Join(v, ",")))
		}
	}

	if len(params.Body) > 0 {
		printlnPrefixLine("-Body:")
		printlnPrefixLine(getBodyByLimited(params.Body))
	}
}

func printHTTPResponse(url string, resp *http.Response) error {
	printDivLine()
	printlnPrefixLine(fmt.Sprintf("-Response: (%d) %s", resp.StatusCode, url))
	printDivLine()

	printlnPrefixLine("-Headers:")
	for k, v := range resp.Header {
		printlnPrefixLine(fmt.Sprintf("%s: %s", k, strings.Join(v, ";")))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		printErrorLine(err)
		return err
	}
	printlnPrefixLine("-Body:")
	printlnPrefixLine(getBodyByLimited(string(body)))

	return nil
}

func getBodyByLimited(body string) string {
	const bodyLenLimit = 1024
	limit := intMin(len(body), bodyLenLimit)
	words := []rune(body)[:limit]
	lines := strings.Split(string(words), "\n")
	return strings.Join(lines, "\n# ")
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func printDivLine() {
	if IsHTTPLog {
		fmt.Println("#", strings.Repeat("*", 60))
	}
}

func printErrorLine(arg interface{}) {
	printlnPrefixLine(fmt.Sprintf("-ERROR: %v", arg))
}

func printlnPrefixLine(arg interface{}) {
	if IsHTTPLog {
		fmt.Println("#", arg)
	}
}

// ********* HTTP Response (server)

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&JSONResponse{Data: m}); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteOKHTMLResp returns html with StatusOK.
func WriteOKHTMLResp(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		WriteErrJSONResp(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// WriteErrJSONResp writes the error response as a Standard API JSON response with a response code.
func WriteErrJSONResp(w http.ResponseWriter, errCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&JSONErrResponse{
		Error: &APIError{Status: errCode, Title: errMsg},
	})
}
