package portout

import (
	"log/slog"
	"net/http"
)

type PortInterface interface {
	Response(w http.ResponseWriter, statusCode int, body any)
	ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, body string)
}

type portImpl struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) PortInterface {
	return &portImpl{
		logger: logger,
	}
}
