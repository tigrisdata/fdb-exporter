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
	"fmt"

	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type WorkloadTransactionsMetricGroup struct {
	MetricGroup
}

func NewWorkloadTransactionsMetricGroup() *WorkloadTransactionsMetricGroup {
	w := WorkloadTransactionsMetricGroup{}
	w.scopes = make(map[string]tally.Scope)
	return &w
}

func (w *WorkloadTransactionsMetricGroup) getValidTagsKeys(scopeName string) []string {
	if scopeName == "default" {
		return []string{
			"env",
			"service",
			"version",
			"cluster",
		}
	} else if scopeName == "started" {
		return []string{
			"env",
			"service",
			"version",
			"cluster",
			"priority",
		}
	} else {
		fmt.Println("unknown scope name")
	}
	return nil
}

func (w *WorkloadTransactionsMetricGroup) getTags(scopeName string, priority string) map[string]string {
	if scopeName == "default" {
		return StandardizeTags(GetBaseTags(), w.getValidTagsKeys(scopeName))
	} else if scopeName == "started" {
		tags := GetBaseTags()
		tags["priority"] = priority
		return StandardizeTags(tags, w.getValidTagsKeys(scopeName))
	} else {
		fmt.Println("unknown scope")
	}
	return nil
}

func (w *WorkloadTransactionsMetricGroup) SetStatus(status *models.FullStatus) {
	w.status = status
}

func (w *WorkloadTransactionsMetricGroup) InitScopes() {
	// two different scopes can have the same name if they need to have different tag set
	w.scopes["default"] = WorkloadScope.SubScope("transactions")
	w.scopes["started"] = WorkloadScope.SubScope("transactions")
}

func (w *WorkloadTransactionsMetricGroup) GetMetrics() {
	metrics := map[string]int{
		"committed":                    w.status.Cluster.Workload.Transactions.Committed.Counter,
		"conflicted":                   w.status.Cluster.Workload.Transactions.Conflicted.Counter,
		"rejected_for_queued_too_long": w.status.Cluster.Workload.Transactions.RejectedForQueuedTooLong.Counter,
	}
	for name, value := range metrics {
		SetIntGauge(w.scopes["default"], name, w.getTags("default", ""), value)
	}

	// The total number of started transactions are started transactions with batch + default + immediate priorities
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "batch"), w.status.Cluster.Workload.Transactions.StartedBatchPriority.Counter)
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "default"), w.status.Cluster.Workload.Transactions.StartedDefaultPriority.Counter)
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "immediate"), w.status.Cluster.Workload.Transactions.StartedImmediatePriority.Counter)
}
