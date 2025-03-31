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

type DataMetricGroup struct {
	metricGroup
}

func NewDataMetricGroup(info *MetricReporter) *DataMetricGroup {
	return &DataMetricGroup{*newMetricGroup("data", info.GetScopeOrExit("cluster"), info)}
}

func (d *DataMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := d.GetScopeOrExit("default")
	if !isValidClusterData(status) || status.Cluster.Data.MovingData == nil {
		log.Error().Msg("failed to get data metric group")
		return
	}

	missingData := 0
	if status.Cluster.Data.State.Name == "missing_data" {
		missingData = 1
	}

	metrics := map[string]interface{}{
		"average_partition_size_bytes":               status.Cluster.Data.AveragePartitionSizeBytes,
		"least_operating_space_bytes_log_server":     status.Cluster.Data.LeastOperatingSpaceBytesLogServer,
		"least_operating_space_bytes_storage_server": status.Cluster.Data.LeastOperatingSpaceBytesStorageServer,
		"moving_data_in_flight_bytes":                status.Cluster.Data.MovingData.InFlightBytes,
		"moving_data_in_queue_bytes":                 status.Cluster.Data.MovingData.InQueueBytes,
		"moving_data_total_written_types":            status.Cluster.Data.MovingData.TotalWrittenBytes,
		"total_disk_used_bytes":                      status.Cluster.Data.TotalDiskUsedBytes,
		"total_kv_size_bytes":                        status.Cluster.Data.TotalKvSizeBytes,
		"min_replicas_remaining":                     status.Cluster.Data.State.MinReplicasRemaining,
		"state_health":                               status.Cluster.Data.State.Healthy,
		"missing_data":                               missingData,
	}
	SetMultipleGauges(scope, metrics, GetBaseTags())
}
