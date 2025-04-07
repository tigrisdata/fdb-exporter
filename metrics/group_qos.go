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

type DbClusterQos struct {
	metricGroup
}

func NewDbClusterQos(info *MetricReporter) *DbClusterQos {
	return &DbClusterQos{*newMetricGroup("qos", info.GetScopeOrExit("cluster"), info)}
}

func (d *DbClusterQos) GetMetrics(status *models.FullStatus) {
	scope := d.GetScopeOrExit("default")
	if !isValidClusterQos(status) {
		log.Error().Msg("failed to get database Qos")
		return
	}
	SetGauge(scope, "transaction_per_second_limit", GetBaseTags(), status.Cluster.Qos.TransactionsPerSecondLimit)
	SetGauge(scope, "released_transactions_per_second", GetBaseTags(), status.Cluster.Qos.ReleasedTransactionsPerSecond)
	SetGauge(scope, "limiting_queue_bytes_storage_server", GetBaseTags(), status.Cluster.Qos.LimitingQueueBytesStorageServer)
	SetGauge(scope, "worst_storage_server_durability_lag_seconds", GetBaseTags(), status.Cluster.Qos.WorstDurabilityLagStorageServer.Seconds)
	SetGauge(scope, "worst_queue_bytes_storage_server", GetBaseTags(), status.Cluster.Qos.WorstQueueBytesStorageServer)
	SetGauge(scope, "worst_queue_bytes_log_server", GetBaseTags(), status.Cluster.Qos.WorstQueueBytesLogServer)
}
