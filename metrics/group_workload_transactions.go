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

func NewWorkloadTransactionsMetricGroup(info *MetricReporter) *WorkloadTransactionsMetricGroup {
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
	transactionsScope := w.GetScopeOrExit("default")
	metrics := make(map[string]interface{})
	if !isValidWorkload(status) {
		log.Error().Msg("failed to get workload metric group")
		return
	}
	if status.Cluster.Workload.Transactions == nil {
		log.Error().Msg("failed to get workload -> transactions metric group")
		return
	}
	workloadTransactions := status.Cluster.Workload.Transactions
	tags := w.getTags("default", "")
	if workloadTransactions.Committed != nil {
		metrics["committed_count"] = workloadTransactions.Committed.Counter
		metrics["committed_hz"] = workloadTransactions.Committed.Hz
	}
	if workloadTransactions.Conflicted != nil {
		metrics["conflicted_count"] = workloadTransactions.Conflicted.Counter
		metrics["conflicted_hz"] = workloadTransactions.Conflicted.Hz
	}
	if workloadTransactions.RejectedForQueuedTooLong != nil {
		metrics["rejected_for_queued_too_long_count"] = workloadTransactions.RejectedForQueuedTooLong.Counter
		metrics["rejected_for_queued_too_long_hz"] = workloadTransactions.RejectedForQueuedTooLong.Hz
	}
	SetMultipleGauges(transactionsScope, metrics, tags)

	// The total number of started transactions are started transactions with batch + default + immediate priorities
	startedScope := w.GetScopeOrExit("started")

	if workloadTransactions.StartedBatchPriority != nil {
		batchMetrics := map[string]interface{}{
			"started_count": workloadTransactions.StartedBatchPriority.Counter,
			"started_hz":    workloadTransactions.StartedBatchPriority.Hz,
		}
		SetMultipleGauges(startedScope, batchMetrics, w.getTags("started", "batch"))
	}

	if workloadTransactions.StartedDefaultPriority != nil {
		defaultMetrics := map[string]interface{}{
			"default_count": workloadTransactions.StartedDefaultPriority.Counter,
			"default_hz":    workloadTransactions.StartedDefaultPriority.Hz,
		}
		SetMultipleGauges(startedScope, defaultMetrics, w.getTags("started", "default"))
	}

	if workloadTransactions.StartedImmediatePriority != nil {
		immediateMetrics := map[string]interface{}{
			"immediate_count": workloadTransactions.StartedImmediatePriority.Counter,
			"immediate_hz":    workloadTransactions.StartedImmediatePriority.Hz,
		}
		SetMultipleGauges(startedScope, immediateMetrics, w.getTags("started", "immediate"))
	}
}
