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
	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/models"
)

type WorkloadKeysMetricGroup struct {
	metricGroup
}

func NewWorkloadKeysMetricGroup(info *MetricInfo) *WorkloadKeysMetricGroup {
	return &WorkloadKeysMetricGroup{*newMetricGroup("keys", info.GetScopeOrExit("workload"), info)}
}

func (w *WorkloadKeysMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := w.GetScopeOrExit("default")
	if status == nil || status.Cluster == nil || status.Cluster.Workload == nil || status.Cluster.Workload.Keys == nil || &status.Cluster.Workload.Keys.Read == nil {
		log.Error().Msg("failed to get workload keys metric group")
		return
	}
	SetIntGauge(scope, "read", GetBaseTags(), status.Cluster.Workload.Keys.Read.Counter)
}
