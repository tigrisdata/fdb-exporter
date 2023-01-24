// Copyright 2022-2023 Tigris Data, Inc.
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
	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/models"
)

type WorkloadOperationsMetricGroup struct {
	metricGroup
}

func NewWorkloadOperationsMetricGroup(info *MetricReporter) *WorkloadOperationsMetricGroup {
	return &WorkloadOperationsMetricGroup{*newMetricGroup("operations", info.GetScopeOrExit("workload"), info)}
}

func (w *WorkloadOperationsMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := w.GetScopeOrExit("default")
	metrics := make(map[string]interface{})
	if !isValidWorkload(status) {
		log.Error().Msg("failed to get workload metric group")
		return
	}
	workloadOperations := status.Cluster.Workload.Operations
	if workloadOperations == nil {
		log.Error().Msg("failed to get workload -> operations")
	}
	if workloadOperations.Reads != nil {
		metrics["reads_count"] = workloadOperations.Reads.Counter
		metrics["reads_hz"] = workloadOperations.Reads.Hz
	}
	if workloadOperations.Writes != nil {
		metrics["writes_count"] = workloadOperations.Writes.Counter
		metrics["writes_hz"] = workloadOperations.Writes.Hz
	}
	if workloadOperations.LocationRequests != nil {
		metrics["location_requests_count"] = workloadOperations.LocationRequests.Counter
		metrics["location_requests_hz"] = workloadOperations.LocationRequests.Hz
	}
	if workloadOperations.LowPriorityReads != nil {
		metrics["low_priority_reads_count"] = workloadOperations.LowPriorityReads.Counter
		metrics["low_priority_reads_hz"] = workloadOperations.LowPriorityReads.Hz
	}
	if workloadOperations.MemoryErrors != nil {
		metrics["memory_errors_count"] = workloadOperations.MemoryErrors.Counter
		metrics["memory_errors_hz"] = workloadOperations.MemoryErrors.Hz
	}
	if workloadOperations.ReadRequests != nil {
		metrics["read_requests_count"] = workloadOperations.ReadRequests.Counter
		metrics["read_requests_hz"] = workloadOperations.ReadRequests.Hz
	}

	SetMultipleGauges(scope, metrics, GetBaseTags())
}
