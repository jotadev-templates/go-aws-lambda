package appcompany

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"crm-lambda/domain/dmncompany"
)

func (i *impl) ExecuteCreate(w http.ResponseWriter, r *http.Request) {
	i.logger.Debug("start ExecuteCreate")

	/*
		Body
	*/
	body, err := io.ReadAll(r.Body)
	if err != nil {
		i.logger.Error(err.Error())
		i.portOut.ResponseError(w, r, http.StatusBadRequest, "body is invalid")
		return
	}
	defer func() {
		if err = r.Body.Close(); err != nil {
			i.logger.Warn(err.Error())
		}
	}()
	i.logger.Debug(fmt.Sprintf("body: %s", string(body)))

	dataBodyDomain, err := dmncompany.ConvertRequestBodyToDomain(body)
	if err != nil {
		i.logger.Error(err.Error())
		i.portOut.ResponseError(w, r, http.StatusBadRequest, "body is invalid")
		return
	}
	i.logger.Debug(fmt.Sprintf("dataBodyDomain: %v", dataBodyDomain))

	/*
		Database
	*/

	if err = i.database.Create(dbTableName, dbTableKey, dataBodyDomain); err != nil {
		isExist := strings.Contains(err.Error(), "ConditionalCheckFailedException")
		if isExist {
			i.portOut.ResponseError(w, r, http.StatusConflict, fmt.Sprintf("ID %s is already exist", dataBodyDomain.ID))
		} else {
			i.logger.Error(err.Error())
			i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		}
		return
	}
	i.portOut.Response(w, http.StatusCreated, nil)
}
