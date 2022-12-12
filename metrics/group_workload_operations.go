package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type WorkloadOperationsMetricGroup struct {
	MetricGroup
}

func NewWorkloadOperationsMetricGroup() *WorkloadOperationsMetricGroup {
	w := WorkloadOperationsMetricGroup{}
	w.scopes = make(map[string]tally.Scope)
	return &w
}

func (w *WorkloadOperationsMetricGroup) SetStatus(status *models.FullStatus) {
	w.status = status
}

func (w *WorkloadOperationsMetricGroup) InitScopes() {
	w.scopes["default"] = WorkloadScope.SubScope("operations")
}

func (w *WorkloadOperationsMetricGroup) GetMetrics() {
	metrics := map[string]int{
		"reads":              w.status.Cluster.Workload.Operations.Reads.Counter,
		"writes":             w.status.Cluster.Workload.Operations.Writes.Counter,
		"location_requests":  w.status.Cluster.Workload.Operations.LocationRequests.Counter,
		"low_priority_reads": w.status.Cluster.Workload.Operations.LowPriorityReads.Counter,
		"memory_errors":      w.status.Cluster.Workload.Operations.MemoryErrors.Counter,
		"read_requests":      w.status.Cluster.Workload.Operations.ReadRequests.Counter,
	}
	for name, value := range metrics {
		SetIntGauge(w.scopes["default"], name, GetBaseTags(), value)
	}
}
