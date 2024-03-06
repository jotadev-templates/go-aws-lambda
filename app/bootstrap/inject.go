package bootstrap

import (
	"log/slog"
	"sync"

	"crm-lambda/adapter/adpaws"
	"crm-lambda/application/appcompany"
	"crm-lambda/port/portout"
)

var (
	mu sync.RWMutex
	i  Inject
)

type Inject struct {
	Logger *slog.Logger
	Port   port
	Adp    adapter
	App    app
}

type port struct {
	Out portout.PortInterface
}
type adapter struct {
	DynamoDB adpaws.DynamoDBInterface
}
type app struct {
	Company appcompany.AppInterface
}

/************************************************
* Constructor
************************************************/

func NewInject() *Inject {
	mu.Lock()
	defer mu.Unlock()

	if i.Logger != nil {
		return &i
	}
	i.Logger = newLogger()

	// Adapter
	i.Adp.DynamoDB = adpaws.NewDynamoDB()

	// Port
	i.Port.Out = portout.New(i.Logger)

	// Application
	i.App.Company = appcompany.New(i.Logger, i.Adp.DynamoDB, i.Port.Out)

	return &i
}
