package appcompany

import (
	"fmt"
	"net/http"

	"crm-lambda/domain/dmncompany"
)

func (i *impl) ExecuteFindByID(w http.ResponseWriter, r *http.Request) {
	i.logger.Debug("start ExecuteFindByID")

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
		Database
	*/
	var outputDB any
	if err = i.database.FindByID(dbTableName, dbTableKey, paramID, &outputDB); err != nil {
		i.logger.Error(err.Error())
		i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		return
	}
	i.logger.Debug(fmt.Sprintf("output DB: %s", outputDB))

	/*
		Data
	*/
	dataDomain, err := dmncompany.ConvertDatabaseToDomain(outputDB)
	if err != nil {
		i.logger.Error(err.Error())
		i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		return
	}
	i.logger.Debug(fmt.Sprintf("dataDomain: %v", dataDomain))

	if len(dataDomain.ID) == 0 {
		i.portOut.ResponseError(w, r, http.StatusNotFound, fmt.Sprintf("ID %s is not found", paramID))
	} else {
		i.portOut.Response(w, http.StatusOK, dataDomain)
	}
}
