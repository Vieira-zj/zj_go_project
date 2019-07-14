package handlers

import (
	"net/http"

	"github.com/golib/httprouter"
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
	// mock demo
	routers = append(routers, RouterEntry{"MockDemoGet", "GET", "/demo/:id", MockDemoHandler})
	routers = append(routers, RouterEntry{"MockDemoPost", "POST", "/demo/:id", MockDemoHandler})
	// mock test
	routers = append(routers, RouterEntry{"MockTestGetPart1", "GET", "/mocktest/one/:id", MockTestHandler01})
	routers = append(routers, RouterEntry{"MockTestGetPart2", "GET", "/mocktest/two/:id", MockTestHandler02})
	// mock qiniu
	routers = append(routers, RouterEntry{"MockQiNiuTestGet", "GET", "/mockqiniu/:id", MockQiNiuHandler})

	router := httprouter.New()
	hooks := NewHooks()
	for _, route := range routers {
		router.Handle(route.Method, route.Path, hooks.RunHooks(route.HandlerFunc))
	}
	router.NotFound = http.HandlerFunc(MockNotFound)
	return router
}
