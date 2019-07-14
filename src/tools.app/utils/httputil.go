package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ********* HTTP Request

var (
	// IsHTTPLog flag print logs for http request and response.
	IsHTTPLog = true
)

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
		log.Println("#", strings.Repeat("*", 60))
	}
}

func printErrorLine(arg interface{}) {
	printlnPrefixLine(fmt.Sprintf("-ERROR: %v", arg))
}

func printlnPrefixLine(arg interface{}) {
	if IsHTTPLog {
		log.Println("#", arg)
	}
}
