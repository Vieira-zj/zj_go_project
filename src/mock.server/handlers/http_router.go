package handlers

import (
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
	routers = append(routers, RouterEntry{"MockDefault", "GET", "/", MockDefault})
	// mock demo
	routers = append(routers, RouterEntry{"MockDemo", "GET", "/demo/:id", MockDemoHandler})
	routers = append(routers, RouterEntry{"MockDemo", "POST", "/demo/:id", MockDemoHandler})
	// mock test
	routers = append(routers, RouterEntry{"MockTestPart1", "GET", "/mocktest/one/:id", MockTestHandler01})
	routers = append(routers, RouterEntry{"MockTestPart2", "GET", "/mocktest/two/:id", MockTestHandler02})
	// mock qiniu
	routers = append(routers, RouterEntry{"MockQiNiuTest", "GET", "/mockqiniu/:id", MockQiNiuHandler})
	// tools
	routers = append(routers, RouterEntry{"Tools", "POST", "/tools/:name", ToolsHandler})

	router := httprouter.New()
	hooks := NewHooks()
	for _, route := range routers {
		router.Handle(route.Method, route.Path, hooks.RunHooks(route.HandlerFunc))
	}
	router.NotFound = WrapHandlerFunc(hooks.RunHooks(MockNotFound))

	return router
}
