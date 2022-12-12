package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type DbStatusMetricGroup struct {
	MetricGroup
}

func NewDbStatusMetricGroup() *DbStatusMetricGroup {
	d := DbStatusMetricGroup{}
	d.scopes = make(map[string]tally.Scope)
	return &d
}

func (d *DbStatusMetricGroup) SetStatus(status *models.FullStatus) {
	d.status = status
}

func (d *DbStatusMetricGroup) InitScopes() {
	d.scopes["default"] = ClientScope.SubScope("status")
}

func (d *DbStatusMetricGroup) GetMetrics() {
	SetBoolGauge(d.scopes["default"], "available", GetBaseTags(), d.status.Client.DatabaseStatus.Available)
	SetBoolGauge(d.scopes["default"], "healthy", GetBaseTags(), d.status.Client.DatabaseStatus.Healthy)
}
