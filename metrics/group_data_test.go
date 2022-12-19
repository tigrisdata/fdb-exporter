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
	"testing"
)

func TestNewDataMetricGroup(t *testing.T) {
	metrics := getMetricsFromTestFile(t, "status-single-basic.json")
	// True represents a non-zero value, false represent zero value
	expected := []string{
		"fdb_cluster_data_average_partition_size_bytes",
		"fdb_cluster_data_least_operating_space_bytes_log_server",
		"fdb_cluster_data_least_operating_space_bytes_storage_server",
		"fdb_cluster_data_moving_data_in_flight_bytes",
		"fdb_cluster_data_moving_data_in_queue_bytes",
		"fdb_cluster_data_moving_data_total_written_types",
		"fdb_cluster_data_total_disk_used_bytes",
		"fdb_cluster_data_total_kv_size_bytes",
	}
	checkMetrics(t, metrics, expected)
}
