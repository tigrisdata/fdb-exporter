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
	metricGroup
}

func NewWorkloadTransactionsMetricGroup(info *MetricInfo) *WorkloadTransactionsMetricGroup {
	parentScope := info.GetScopeOrExit("workload")
	w := WorkloadTransactionsMetricGroup{*newMetricGroup("transactions", parentScope, info)}
	w.AddScope(parentScope, "started", "transactions")
	return &w
}

func (w *WorkloadTransactionsMetricGroup) getValidTagsKeys(scopeKey string) []string {
	if scopeKey == "started" {
		return append(getBaseTagKeys(), "priority")
	} else {
		return getBaseTagKeys()
	}
}

func (w *WorkloadTransactionsMetricGroup) getTags(scopeKey string, priority string) map[string]string {
	if scopeKey == "default" {
		return StandardizeTags(GetBaseTags(), w.getValidTagsKeys(scopeKey))
	} else if scopeKey == "started" {
		tags := GetBaseTags()
		tags["priority"] = priority
		return StandardizeTags(tags, w.getValidTagsKeys(scopeKey))
	} else {
		log.Error().Msg("unknown scope")
	}
	return nil
}

func (w *WorkloadTransactionsMetricGroup) GetMetrics(status *models.FullStatus) {
	transActionsScope := w.GetScopeOrExit("default")
	tags := w.getTags("default", "")
	metrics := map[string]int{
		"committed":                    status.Cluster.Workload.Transactions.Committed.Counter,
		"conflicted":                   status.Cluster.Workload.Transactions.Conflicted.Counter,
		"rejected_for_queued_too_long": status.Cluster.Workload.Transactions.RejectedForQueuedTooLong.Counter,
	}
	for name, value := range metrics {
		SetIntGauge(transActionsScope, name, tags, value)
	}

	// The total number of started transactions are started transactions with batch + default + immediate priorities
	startedScope := w.GetScopeOrExit("started")
	SetIntGauge(startedScope, "started", w.getTags("started", "batch"), status.Cluster.Workload.Transactions.StartedBatchPriority.Counter)
	SetIntGauge(startedScope, "started", w.getTags("started", "default"), status.Cluster.Workload.Transactions.StartedDefaultPriority.Counter)
	SetIntGauge(startedScope, "started", w.getTags("started", "immediate"), status.Cluster.Workload.Transactions.StartedImmediatePriority.Counter)
}
