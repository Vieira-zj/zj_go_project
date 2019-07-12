package handlers

import (
	"github.com/golib/httprouter"
	"mock.server/common"
)

// RouterEntry an router entry.
type RouterEntry struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

// NewHTTPRouter returns a new http server.
func NewHTTPRouter() *httprouter.Router {
	routers := make([]RouterEntry, 0, 10)
	routers = append(routers, RouterEntry{"MockDemoGet", "GET", "/demo/:id", MockDemoHandler})
	routers = append(routers, RouterEntry{"MockDemoPost", "POST", "/demo/:id", MockDemoHandler})

	router := httprouter.New()
	for _, route := range routers {
		router.Handle(route.Method, route.Path, common.PerfLogger(route.HandlerFunc))
	}
	return router
}
