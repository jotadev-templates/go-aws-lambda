package appcompany

import (
	"fmt"
	"net/http"
	"strings"
)

type executeCommonImpl struct {
	impl
}

func (i *executeCommonImpl) getParamID(w http.ResponseWriter, r *http.Request) (string, error) {
	var result string
	errReturn := fmt.Errorf("error")

	urlPath := r.URL.Path
	splitURL := strings.Split(urlPath, "/")

	if len(splitURL) != 3 {
		i.logger.Error(fmt.Sprintf("invalid url path: %s", urlPath))
		i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		return "", errReturn
	}
	result = splitURL[2]

	if len(strings.TrimSpace(result)) == 0 {
		i.portOut.ResponseError(w, r, http.StatusBadRequest, "ID is empty")
		return "", errReturn
	}
	return result, nil
}
