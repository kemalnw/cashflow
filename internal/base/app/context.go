package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
	Request   *http.Request
	AppConfig *Config
}

func NewContext(c *gin.Context, appConfig *Config) *Context {
	return &Context{
		Context:   c,
		Request:   c.Request,
		AppConfig: appConfig,
	}
}
