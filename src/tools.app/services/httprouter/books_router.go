package services

import (
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// RouterEntry an router entry.
type RouterEntry struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// BooksServer contains router entries.
type BooksServer struct {
	routers []RouterEntry
}

// NewBooksSvr retruns a http server instance.
func NewBooksSvr() *BooksServer {
	handler := NewBooksHandler()
	routers := []RouterEntry{
		RouterEntry{"Index", "GET", "/", handler.Index},
		RouterEntry{"BookIndex", "GET", "/books", handler.BookIndex},
		RouterEntry{"Bookshow", "GET", "/books/:isdn", handler.BookShow},
		RouterEntry{"Bookshow", "POST", "/books", handler.BookCreate},
	}
	return &BooksServer{routers: routers}
}

// GetHTTPRouter retruns a httpr router with handlers.
func (svc BooksServer) GetHTTPRouter() *httprouter.Router {
	router := httprouter.New()
	for _, route := range svc.routers {
		router.Handle(route.Method, route.Path, svc.logger(route.HandlerFunc))
	}
	return router
}

func (svc BooksServer) logger(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log.Printf("start: %s %s\n", r.Method, r.URL.Path)
		fn(w, r, param)
		log.Printf("done: %v (%s %s)\n", time.Since(start), r.Method, r.URL.Path)
	}
}
