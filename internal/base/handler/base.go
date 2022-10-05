package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kemalnw/cashflow/internal/base/app"
	"github.com/kemalnw/cashflow/pkg/server"
	"gorm.io/gorm"
)

type HandlerFn func(ctx *app.Context) *server.Response

type BaseHTTPHandler struct {
	Handlers  interface{}
	Params    map[string]string
	appConfig *app.Config
	DB        *gorm.DB
}

func (b *BaseHTTPHandler) RunAction(handler HandlerFn) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := handler(app.NewContext(c, b.appConfig))
		httpStatus := resp.GetStatus()

		c.JSON(httpStatus, resp)
		return
	}
}

func (b *BaseHTTPHandler) AsJson(status int, message string, data interface{}) *server.Response {
	return &server.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewBaseHandler(appConfig *app.Config, db *gorm.DB) *BaseHTTPHandler {
	return &BaseHTTPHandler{
		appConfig: appConfig,
		DB:        db,
	}
}
