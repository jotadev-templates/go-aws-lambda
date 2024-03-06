package portout

import (
	"encoding/json"
	"net/http"
	"strings"
)

const (
	contentType     string = "Content-Type"
	applicationJSON string = "application/json"
)

const (
	msgErrStatusCodeIsEmpty string = "status code is empty"
)

func (i *portImpl) Response(w http.ResponseWriter, statusCode int, body any) {
	if statusCode == 0 {
		i.logger.Warn(msgErrStatusCodeIsEmpty)
	} else {
		w.WriteHeader(statusCode)
	}
	w.Header().Set(contentType, applicationJSON)

	if body != nil {
		if err := json.NewEncoder(w).Encode(body); err != nil {
			i.logger.Error(err.Error())
		}
	}
}

func (i *portImpl) ResponseError(w http.ResponseWriter, r *http.Request, statusCode int, body string) {
	if statusCode == 0 {
		i.logger.Warn(msgErrStatusCodeIsEmpty)
	} else {
		w.WriteHeader(statusCode)
	}
	w.Header().Set(contentType, applicationJSON)

	type response struct {
		StatusCode int    `json:"statusCode,omitempty"`
		Method     string `json:"method"`
		Body       string `json:"body,omitempty"`
	}
	resp := response{
		StatusCode: statusCode,
		Method:     r.Method,
		Body:       body,
	}

	if len(strings.TrimSpace(body)) > 0 {
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			i.logger.Error(err.Error())
		}
	}
}
