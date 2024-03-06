package portin

import (
	"net/http"

	"crm-lambda/bootstrap"
)

type route struct {
	endpoint    string
	executeFunc func(w http.ResponseWriter, r *http.Request)
}

// ****************************************

func NewHandler(i *bootstrap.Inject) {
	for _, r := range buildRoutes(i) {
		finalHandler := http.HandlerFunc(r.executeFunc)
		http.Handle(r.endpoint, middlewareRequestID(i.Logger, finalHandler))
	}
}
