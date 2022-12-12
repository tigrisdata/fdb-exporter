// Copyright 2022 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
