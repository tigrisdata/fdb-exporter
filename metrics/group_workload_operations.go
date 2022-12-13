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

import "github.com/tigrisdata/fdb-exporter/models"

type WorkloadOperationsMetricGroup struct {
	MetricGroup
}

func NewWorkloadOperationsMetricGroup(mInfo *MetricInfo) *WorkloadOperationsMetricGroup {
	return &WorkloadOperationsMetricGroup{*NewMetricGroup("operations", mInfo.scopes["workload"], mInfo)}
}

func (w *WorkloadOperationsMetricGroup) GetMetrics(status *models.FullStatus) {
	metrics := map[string]int{
		"reads":              status.Cluster.Workload.Operations.Reads.Counter,
		"writes":             status.Cluster.Workload.Operations.Writes.Counter,
		"location_requests":  status.Cluster.Workload.Operations.LocationRequests.Counter,
		"low_priority_reads": status.Cluster.Workload.Operations.LowPriorityReads.Counter,
		"memory_errors":      status.Cluster.Workload.Operations.MemoryErrors.Counter,
		"read_requests":      status.Cluster.Workload.Operations.ReadRequests.Counter,
	}
	for name, value := range metrics {
		SetIntGauge(w.scopes["default"], name, GetBaseTags(), value)
	}
}
