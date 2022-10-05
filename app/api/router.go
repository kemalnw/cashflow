package api

import (
	"fmt"
	"github.com/kemalnw/cashflow/internal/base/handler"
)

func (h *HttpServe) setupRouter() {
	h.Route("GET", "/", h.base.HealthCheck)
}

func (h *HttpServe) Route(method string, path string, f handler.HandlerFn) {
	switch method {
	case "GET":
		h.router.GET(path, h.base.RunAction(f))
	case "POST":
		h.router.POST(path, h.base.RunAction(f))
	case "PUT":
		h.router.PUT(path, h.base.RunAction(f))
	case "DELETE":
		h.router.DELETE(path, h.base.RunAction(f))
	default:
		panic(fmt.Sprintf(":%s method not allow", method))
	}
}
