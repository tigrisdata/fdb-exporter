package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type WorkloadBytesMetricGroup struct {
	MetricGroup
}

func NewWorkloadBytesMetricGroup() *WorkloadBytesMetricGroup {
	w := WorkloadBytesMetricGroup{}
	w.scopes = make(map[string]tally.Scope)
	return &w
}

func (w *WorkloadBytesMetricGroup) SetStatus(status *models.FullStatus) {
	w.status = status
}

func (w *WorkloadBytesMetricGroup) InitScopes() {
	w.scopes["default"] = WorkloadScope.SubScope("keys")
}

func (w *WorkloadBytesMetricGroup) GetMetrics() {
	SetIntGauge(w.scopes["default"], "read", GetBaseTags(), w.status.Cluster.Workload.Bytes.Read.Counter)
	SetIntGauge(w.scopes["default"], "written", GetBaseTags(), w.status.Cluster.Workload.Bytes.Written.Counter)
}
