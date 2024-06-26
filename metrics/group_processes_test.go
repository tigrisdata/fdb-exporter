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
	"fmt"
	"strings"
	"testing"
)

func TestProcessesMetricGroupSingleBasic(t *testing.T) {
	initTestMetricReporter()
	metrics := getMetricsFromTestFile(t, "status-single-basic.json")
	// True represents a non-zero value, false represent zero value
	expected := []string{
		"fdb_cluster_processes_cpu_cores",
		"fdb_cluster_processes_disk_busy",
		"fdb_cluster_processes_disk_total_bytes",
		"fdb_cluster_processes_disk_reads_count",
		"fdb_cluster_processes_disk_reads_hz",
		"fdb_cluster_processes_disk_writes_count",
		"fdb_cluster_processes_disk_writes_hz",
		"fdb_cluster_processes_mem_available_bytes",
		"fdb_cluster_processes_mem_limit_bytes",
		"fdb_cluster_processes_mem_unused_allocated_memory",
		"fdb_cluster_processes_mem_unused_bytes",
		"fdb_cluster_processes_network_conn_errors_hz",
		"fdb_cluster_processes_network_conn_closed_hz",
		"fdb_cluster_processes_network_conn_established",
		"fdb_cluster_processes_network_current_connections",
		"fdb_cluster_processes_network_megabits_sent",
		"fdb_cluster_processes_network_megabits_received",
		"fdb_cluster_grv_latency",
		"fdb_cluster_commit_latency",
		"fdb_cluster_processes_kvstore_available_bytes",
		"fdb_cluster_processes_kvstore_free_bytes",
		"fdb_cluster_processes_kvstore_total_bytes",
		"fdb_cluster_processes_kvstore_used_bytes",
		"fdb_cluster_processes_stored_bytes",
		"fdb_cluster_processes_queue_disk_available_bytes",
		"fdb_cluster_processes_queue_disk_free_bytes",
		"fdb_cluster_processes_queue_disk_total_bytes",
		"fdb_cluster_processes_queue_disk_used_bytes",
		"fdb_cluster_processes_data_lag_seconds",
		"fdb_cluster_processes_data_lag_versions",
		"fdb_cluster_processes_durability_lag_seconds",
		"fdb_cluster_processes_durability_lag_versions",
		"fdb_cluster_processes_log_queue_length",
		"fdb_cluster_processes_input_bytes",
		"fdb_cluster_processes_durability_bytes",
	}
	checkMetrics(t, metrics, expected)
}

func TestProcessesMetricGroupMessages(t *testing.T) {
	initTestMetricReporter()
	metrics := getMetricsFromTestFile(t, "status-process-io-timeout.json")
	for _, metric := range metrics {
		if strings.Contains(metric.key, "message") {
			fmt.Println(metric.key)
		}

	}
	expected := []string{
		"fdb_cluster_per_process_messages",
	}
	checkMetrics(t, metrics, expected)
}
