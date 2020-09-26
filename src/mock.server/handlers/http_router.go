package handlers

import (
	"src/mock.server/common"

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

	// Issue: router matcher do not support regexp
	if !common.IsProd() {
		routers = append(routers, RouterEntry{"MockDefault", "OPTIONS", "/ping", MockDefault})
		// mock api
		routers = append(routers, RouterEntry{"MockAPIRegister", "OPTIONS", "/mock/register/:uri", MockAPIRegisterHandler})
		routers = append(routers, RouterEntry{"MockAPI", "OPTIONS", "/mock/api/:uri", MockAPIHandler})
	}

	routers = append(routers, RouterEntry{"MockDefault", "GET", "/ping", MockDefault})
	// mock api
	routers = append(routers, RouterEntry{"MockAPIRegister", "POST", "/mock/register/:uri", MockAPIRegisterHandler})
	routers = append(routers, RouterEntry{"MockAPI", "GET", "/mock/api/:uri", MockAPIHandler})
	routers = append(routers, RouterEntry{"MockAPI", "POST", "/mock/api/:uri", MockAPIHandler})

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
