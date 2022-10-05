package handler

import (
	"github.com/kemalnw/cashflow/internal/base/app"
	"github.com/kemalnw/cashflow/pkg/server"
	"net/http"
)

func (b *BaseHTTPHandler) HealthCheck(_ *app.Context) *server.Response {
	db, err := b.DB.DB()
	if err != nil {
		return b.AsJson(http.StatusInternalServerError, "Error setup DB.", nil)
	}

	err = db.Ping()
	if err != nil {
		return b.AsJson(http.StatusInternalServerError, "Database service is stopped.", nil)
	}

	return b.AsJson(http.StatusOK, "Service is running.", nil)
}
