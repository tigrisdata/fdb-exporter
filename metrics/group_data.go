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

// TODO: finish this
type DataMetricGroup struct {
	MetricGroup
}

func NewDataMetricGroup() *DataMetricGroup {
	d := DataMetricGroup{}
	d.scopes = make(map[string]tally.Scope)
	return &d
}

func (d *DataMetricGroup) SetStatus(status *models.FullStatus) {
	d.status = status
}

func (d *DataMetricGroup) InitScopes() {
	d.scopes["default"] = ClusterScope.SubScope("data")
}

func (d *DataMetricGroup) GetMetrics() {
	metrics := map[string]int{
		"average_partition_size_bytes":               d.status.Cluster.Data.AveragePartitionSizeBytes,
		"least_operating_space_bytes_log_server":     d.status.Cluster.Data.LeastOperatingSpaceBytesLogServer,
		"least_operating_space_bytes_storage_server": d.status.Cluster.Data.LeastOperatingSpaceBytesStorageServer,
		"moving_data_in_flight_bytes":                d.status.Cluster.Data.MovingData.InFlightBytes,
		"moving_data_in_queue_bytes":                 d.status.Cluster.Data.MovingData.InQueueBytes,
		"moving_data_total_written_types":            d.status.Cluster.Data.MovingData.TotalWrittenBytes,
		"total_disk_used_bytes":                      d.status.Cluster.Data.TotalDiskUsedBytes,
		"total_kv_size_bytes":                        d.status.Cluster.Data.TotalKvSizeBytes,
	}
	for name, value := range metrics {
		SetIntGauge(d.scopes["default"], name, GetBaseTags(), value)
	}
}
