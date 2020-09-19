package httprouter_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mysvc "src/tools.app/services/httprouter"

	"github.com/julienschmidt/httprouter"
)

func TestBookShow(t *testing.T) {
	handler := mysvc.NewBooksHandler()

	t.Log("Case1: When the books' isdn does not exist")
	req, err := http.NewRequest("GET", "/books/1234", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := newRequestRecorder(req, "GET", "/books/:isdn", handler.BookShow)
	if rr.Code != http.StatusNotFound {
		t.Error("Expected response code to be 404")
	}
	jsonResp := `{"error":{"status":404,"title":"Book Record Not Found!"}}`
	if strings.Trim(rr.Body.String(), "\n") != jsonResp {
		t.Error("Response body does not match")
	}

	t.Log("Case2: When the book exists")
	req, err = http.NewRequest("GET", "/books/002", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = newRequestRecorder(req, "GET", "/books/:isdn", handler.BookShow)
	if rr.Code != http.StatusOK {
		t.Error("Expected response code to be 200")
	}
	jsonResp = `{"meta":null,"data":{"isdn":"002","title":"To Kill a Mocking Bird","author":"Harper Lee","pages":320}}`
	if strings.Trim(rr.Body.String(), "\n") != jsonResp {
		t.Error("Response body does not match")
	}
}

// Mocks a handler and returns a httptest.ResponseRecorder
func newRequestRecorder(req *http.Request, method, strPath string, fnHandler httprouter.Handle) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
