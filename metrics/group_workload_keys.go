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

type WorkloadKeysMetricGroup struct {
	metricGroup
}

func NewWorkloadKeysMetricGroup(info *MetricReporter) *WorkloadKeysMetricGroup {
	return &WorkloadKeysMetricGroup{*newMetricGroup("keys", info.GetScopeOrExit("workload"), info)}
}

func (w *WorkloadKeysMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := w.GetScopeOrExit("default")
	metrics := make(map[string]interface{})
	if !isValidWorkload(status) {
		log.Error().Msg("failed to get workload keys metric group")
		return
	}
	workloadKeys := status.Cluster.Workload.Keys
	if workloadKeys.Read != nil {
		metrics["read_count"] = workloadKeys.Read.Counter
		metrics["read_hz"] = workloadKeys.Read.Hz
	}
	SetMultipleGauges(scope, metrics, GetBaseTags())
}
