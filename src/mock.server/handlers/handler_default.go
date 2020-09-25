package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"src/mock.server/common"

	"github.com/golib/httprouter"
)

var count int

// MockDefault sends a default message.
func MockDefault(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	count = count + 1
	msg := fmt.Sprintf("Mock Server, access count: %d", count)
	if err := common.WriteOKHTMLResp(w, []byte(msg)); err != nil {
		common.ErrHandler(w, err)
	}
}

// MockNotFound sends a server not found page.
func MockNotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ret := fmt.Sprintf("Page not found for path: %s", r.RequestURI)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(ret)))
	w.WriteHeader(http.StatusNotFound)
	log.Printf("Page not found: %s\n", r.URL.Path)

	if _, err := w.Write([]byte(ret)); err != nil {
		common.ErrHandler(w, err)
	}
}
