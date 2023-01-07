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

func TestLatencyProbeMetricGroupSingleBasic(t *testing.T) {
	initTestMetricReporter()
	metrics := getMetricsFromTestFile(t, "status-single-basic.json")
	// True represents a non-zero value, false represent zero value
	expected := []string{
		"fdb_cluster_latency_probe_commit_seconds",
		"fdb_cluster_latency_probe_read_seconds",
		"fdb_cluster_latency_probe_transaction_start_seconds",
		"fdb_cluster_latency_probe_transaction_start_seconds_batch",
		"fdb_cluster_latency_probe_transaction_start_seconds_immediate",
	}
	checkMetrics(t, metrics, expected)
}
