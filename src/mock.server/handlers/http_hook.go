package handlers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/golib/httprouter"
	"mock.server/common"
)

// ******** Http Connect Hooks

// NewHooks returns http connect handler hooks.
func NewHooks() *Hooks {
	return &Hooks{start: time.Now()}
}

// Hooks http connect handler hooks.
type Hooks struct {
	start time.Time
	mutex sync.Mutex
}

// RunHooks run before and after hooks when handle http connect.
func (h *Hooks) RunHooks(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		defer func() {
			if p := recover(); p != nil {
				common.ErrHandler(w, p.(error))
			}
		}()

		h.mutex.Lock()
		defer func() {
			h.mutex.Unlock()
		}()

		h.beforeHooks(w, r)
		fn(w, r, param)
		h.afterHooks(w, r)
	}
}

func (h *Hooks) beforeHooks(w http.ResponseWriter, r *http.Request) {
	common.LogDivLine()
	log.Printf("Start: %s %s\n", r.Method, r.URL.Path)

	if err := common.LogRequestData(r); err != nil {
		common.ErrHandler(w, err)
		return
	}
	h.start = time.Now()
}

func (h *Hooks) afterHooks(w http.ResponseWriter, r *http.Request) {
	log.Printf("Done (%s %s): %v\n", r.Method, r.URL.Path, time.Since(h.start))
	common.LogDivLine()
}
