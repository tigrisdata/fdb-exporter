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

import "testing"

func TestWorkloadTransactionsMetricGroupSingleBasic(t *testing.T) {
	initTestMetricReporter()
	metrics := getMetricsFromTestFile(t, "status-single-basic.json")
	// True represents a non-zero value, false represent zero value
	expected := []string{
		"fdb_cluster_workload_transactions_committed_count",
		"fdb_cluster_workload_transactions_committed_hz",
		"fdb_cluster_workload_transactions_conflicted_count",
		"fdb_cluster_workload_transactions_conflicted_hz",
		"fdb_cluster_workload_transactions_rejected_for_queued_too_long_count",
		"fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz",
		"fdb_cluster_workload_transactions_started_count",
		"fdb_cluster_workload_transactions_started_hz",
	}
	checkMetrics(t, metrics, expected)
	startedTags := []string{
		"priority",
	}
	checkTagsForMetric(t, metrics, "fdb_cluster_workload_transactions_started_count", startedTags)
	checkTagsForMetric(t, metrics, "fdb_cluster_workload_transactions_started_hz", startedTags)
}
