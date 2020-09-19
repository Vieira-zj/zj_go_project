package utils_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	myutils "src/tools.app/utils"
)

func TestHTTPGet(t *testing.T) {
	const tURL = "http://127.0.0.1:17891/index"
	t.Log(fmt.Sprintf("Case01: http get request test: '%s'", tURL))

	headers := make(http.Header)
	headers.Set("X-Tag", "Http-Get-Test")

	var query url.Values = make(url.Values)
	query.Set("id", "001")

	req := &myutils.HTTPRequest{
		URL:     tURL,
		Headers: headers,
		Query:   query,
	}

	resp, err := myutils.HTTPGet(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Failed: returned code is not OK (%d)", resp.StatusCode)
	}
}

func TestHTTPPost(t *testing.T) {
	const tURL = "http://127.0.0.1:17891/index"
	t.Log(fmt.Sprintf("Case01: http post request test: '%s'", tURL))
	// myutils.IsHTTPLog = false

	headers := make(http.Header)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("X-Tag", "Http-Post-Test")

	req := &myutils.HTTPRequest{
		URL:     tURL,
		Headers: headers,
		Body:    "key1=value1;key2=value2",
	}

	resp, err := myutils.HTTPPost(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Failed: returned code is not OK (%d)", resp.StatusCode)
	}
}
