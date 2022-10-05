package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kemalnw/cashflow/internal/base/handler"
	"github.com/kemalnw/cashflow/pkg/server"
	"os"
)

type HttpServe struct {
	router *gin.Engine
	base   *handler.BaseHTTPHandler
}

func (h *HttpServe) Run() error {
	h.setupRouter()
	h.base.Handlers = h

	return h.router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_SERVER_PORT")))
}

func New(base *handler.BaseHTTPHandler) server.App {
	r := gin.New()

	return &HttpServe{
		base:   base,
		router: r,
	}
}
