package appcompany

import (
	"fmt"
	"net/http"
	"strings"
)

func (i *impl) ExecuteDelete(w http.ResponseWriter, r *http.Request) {
	i.logger.Debug("start ExecuteDelete")

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
	if err = i.database.Delete(dbTableName, dbTableKey, paramID); err != nil {
		isNotExist := strings.Contains(err.Error(), "ConditionalCheckFailedException")
		if isNotExist {
			i.portOut.ResponseError(w, r, http.StatusNotFound, fmt.Sprintf("ID %s is not found", paramID))
		} else {
			i.logger.Error(err.Error())
			i.portOut.ResponseError(w, r, http.StatusInternalServerError, "")
		}
		return
	}
	i.portOut.Response(w, http.StatusNoContent, nil)
}
