package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

// TODO: Metric group and metric interfaces, each one will plug into these
type MetricGroup struct {
	scopes map[string]tally.Scope
	status *models.FullStatus
}

type Collectable interface {
	SetStatus(status *models.FullStatus)
	InitScopes()
	GetMetrics()
}
