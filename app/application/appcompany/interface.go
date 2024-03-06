package appcompany

import (
	"log/slog"
	"net/http"

	"crm-lambda/adapter/adpaws"
	"crm-lambda/port/portout"
)

type AppInterface interface {
	ExecuteFindByID(w http.ResponseWriter, r *http.Request)
	ExecuteCreate(w http.ResponseWriter, r *http.Request)
	ExecuteUpsertAll(w http.ResponseWriter, r *http.Request)
	ExecuteDelete(w http.ResponseWriter, r *http.Request)
}

type impl struct {
	logger   *slog.Logger
	database adpaws.DynamoDBInterface
	portOut  portout.PortInterface
}

func New(
	logger *slog.Logger,
	database adpaws.DynamoDBInterface,
	portOut portout.PortInterface,
) AppInterface {
	return &impl{
		logger:   logger,
		database: database,
		portOut:  portOut,
	}
}
