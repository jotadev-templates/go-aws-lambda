package appcompany

import (
	"fmt"
	"io"
	"net/http"

	"crm-lambda/domain/dmncompany"
)

func (i *impl) ExecuteUpsertAll(w http.ResponseWriter, r *http.Request) {
	i.logger.Debug("start ExecuteUpsertAll")

	var iCommon executeCommonImpl

	/*
		Param
	*/
	paramID, err := iCommon.getParamID(w, r)
	if err != nil {
		return
	}
	i.logger.Debug(fmt.Sprintf("paramID: %s", paramID))

	/*
		Body
	*/
	body, err := io.ReadAll(r.Body)
	if err != nil {
		i.logger.Debug(err.Error())
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
		i.logger.Debug(err.Error())
		i.portOut.ResponseError(w, r, http.StatusBadRequest, "body is invalid")
		return
	}
	i.logger.Debug(fmt.Sprintf("dataBodyDomain: %v", dataBodyDomain))

	isIDValid := paramID == dataBodyDomain.ID
	if !isIDValid {
		i.portOut.ResponseError(w, r, http.StatusBadRequest, fmt.Sprintf(
			"ID are inconsistent: %s and %s are different", paramID, dataBodyDomain.ID))
		return
	}

	/*
		Database
	*/
	if err = i.database.UpsertAll(dbTableName, dataBodyDomain); err != nil {
		i.logger.Error(err.Error())
		i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		return
	}
	i.portOut.Response(w, http.StatusOK, nil)
}
