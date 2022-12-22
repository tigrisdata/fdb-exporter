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

import "testing"

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
	}
	checkMetrics(t, metrics, expected)
}
