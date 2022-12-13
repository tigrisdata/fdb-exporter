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

type WorkloadTransactionsMetricGroup struct {
	MetricGroup
}

func NewWorkloadTransactionsMetricGroup(mInfo *MetricInfo) *WorkloadTransactionsMetricGroup {
	w := WorkloadTransactionsMetricGroup{*NewMetricGroup("transactions", mInfo.scopes["workload"], mInfo)}
	w.AddScope(mInfo.scopes["workload"], "started")
	return &w
}

func (w *WorkloadTransactionsMetricGroup) getValidTagsKeys(scopeName string) []string {
	if scopeName == "started" {
		return append(getBaseTagKeys(), "priority")
	} else {
		return getBaseTagKeys()
	}
}

func (w *WorkloadTransactionsMetricGroup) getTags(scopeName string, priority string) map[string]string {
	if scopeName == "default" {
		return StandardizeTags(GetBaseTags(), w.getValidTagsKeys(scopeName))
	} else if scopeName == "started" {
		tags := GetBaseTags()
		tags["priority"] = priority
		return StandardizeTags(tags, w.getValidTagsKeys(scopeName))
	} else {
		log.Error().Msg("unknown scope")
	}
	return nil
}

func (w *WorkloadTransactionsMetricGroup) GetMetrics(status *models.FullStatus) {
	metrics := map[string]int{
		"committed":                    status.Cluster.Workload.Transactions.Committed.Counter,
		"conflicted":                   status.Cluster.Workload.Transactions.Conflicted.Counter,
		"rejected_for_queued_too_long": status.Cluster.Workload.Transactions.RejectedForQueuedTooLong.Counter,
	}
	for name, value := range metrics {
		SetIntGauge(w.scopes["default"], name, w.getTags("default", ""), value)
	}

	// The total number of started transactions are started transactions with batch + default + immediate priorities
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "batch"), status.Cluster.Workload.Transactions.StartedBatchPriority.Counter)
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "default"), status.Cluster.Workload.Transactions.StartedDefaultPriority.Counter)
	SetIntGauge(w.scopes["started"], "started", w.getTags("started", "immediate"), status.Cluster.Workload.Transactions.StartedImmediatePriority.Counter)
}
