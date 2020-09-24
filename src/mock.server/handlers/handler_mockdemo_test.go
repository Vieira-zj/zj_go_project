package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"

	"src/mock.server/common"
	"src/mock.server/handlers"

	"github.com/golib/httprouter"
)

func TestMockDemo01(t *testing.T) {
	t.Log("Case01: test mock demo01, get resquest.")
	method := "GET"
	req, err := http.NewRequest(method, "/demo/1?userid=xxx&username=xxx", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := newRequestRecorder(req, method, "/demo/:id", handlers.MockDemoHandler)
	if rr.Code != http.StatusOK {
		t.Error("Unexpected returned code:", rr.Code)
	}
	t.Log("Response:", rr.Body.String())
}

func TestMockDemo02(t *testing.T) {
	t.Log("Case01: test mock demo02, post resquest with cookie.")
	method := "POST"
	req, err := http.NewRequest(method, "/demo/4", strings.NewReader("key1=val1;key2=val2"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(common.TextContentType, common.ContentTypeForm)
	req.Header.Set("Cookie", "user=user_001;pwd=test_com")

	rr := newRequestRecorder(req, method, "/demo/:id", handlers.MockDemoHandler)
	if rr.Code != http.StatusOK {
		t.Error("Unexpected returned code:", rr.Code)
	}
	t.Log("Response:", rr.Body.String())
}

// Mocks a handler and returns a httptest.ResponseRecorder.
func newRequestRecorder(req *http.Request, method, strPath string, fnHandler httprouter.Handle) *httptest.ResponseRecorder {
	router := httprouter.New()
	router.Handle(method, strPath, fnHandler)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}
