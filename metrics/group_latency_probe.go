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

type LatencyProbeMetricGroup struct {
	metricGroup
}

func NewLatencyProbeMetricGroup(info *MetricReporter) *LatencyProbeMetricGroup {
	return &LatencyProbeMetricGroup{*newMetricGroup("latency_probe", info.GetScopeOrExit("cluster"), info)}
}

func (d *LatencyProbeMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := d.GetScopeOrExit("default")
	if !isValidClusterLatencyProbe(status) {
		log.Error().Msg("failed to get cluster latency probes")
		return
	}
	latencyProbe := status.Cluster.LatencyProbe
	SetGauge(scope, "commit_seconds", GetBaseTags(), latencyProbe.CommitSeconds)
	SetGauge(scope, "read_seconds", GetBaseTags(), latencyProbe.ReadSeconds)
	SetGauge(scope, "transaction_start_seconds", GetBaseTags(), latencyProbe.TransactionStartSeconds)
	SetGauge(scope, "transaction_start_seconds_immediate", GetBaseTags(), latencyProbe.ImmediatePriorityTransactionStartSeconds)
	SetGauge(scope, "transaction_start_seconds_batch", GetBaseTags(), latencyProbe.BatchPriorityTransactionStartSeconds)
}
