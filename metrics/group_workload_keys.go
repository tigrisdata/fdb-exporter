package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type WorkloadKeysMetricGroup struct {
	MetricGroup
}

func NewWorkloadKeysMetricGroup() *WorkloadKeysMetricGroup {
	w := WorkloadKeysMetricGroup{}
	w.scopes = make(map[string]tally.Scope)
	return &w
}

func (w *WorkloadKeysMetricGroup) getValidTagKeys(_ string) []string {
	return []string{
		"env",
		"service",
		"version",
		"cluster",
	}
}

func (w *WorkloadKeysMetricGroup) getTags(_ string) map[string]string {
	return StandardizeTags(GetBaseTags(), w.getValidTagKeys("default"))
}

func (w *WorkloadKeysMetricGroup) SetStatus(status *models.FullStatus) {
	w.status = status
}

func (w *WorkloadKeysMetricGroup) InitScopes() {
	w.scopes["default"] = WorkloadScope.SubScope("keys")
}

func (w *WorkloadKeysMetricGroup) GetMetrics() {
	SetIntGauge(w.scopes["default"], "read", w.getTags("default"), w.status.Cluster.Workload.Keys.Read.Counter)
}
