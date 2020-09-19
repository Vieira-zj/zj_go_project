package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"src/mock.server/common"

	"github.com/golib/httprouter"
)

// MockDefault sends a mock default page.
func MockDefault(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := common.WriteOKHTMLResp(w, []byte("Mock Default Page")); err != nil {
		common.ErrHandler(w, err)
	}
}

// MockNotFound sends a server not found page.
func MockNotFound(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ret := fmt.Sprintf("Default Not Found Page.\nPage not found for path: %s\n", r.RequestURI)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(ret)))
	w.WriteHeader(http.StatusNotFound)
	log.Printf("Page not found: %s\n", r.URL.Path)

	if _, err := w.Write([]byte(ret)); err != nil {
		common.ErrHandler(w, err)
	}
}
