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

package models

type Workload struct {
	Bytes        *WorkloadBytes        `json:"bytes"`
	Keys         *WorkloadKeys         `json:"keys"`
	Operations   *WorkloadOperations   `json:"operations"`
	Transactions *WorkloadTransactions `json:"transactions"`
}

type WorkloadMetrics struct {
	Counter   int64   `json:"counter"`
	Hz        float64 `json:"hz"`
	Roughness float64 `json:"roughness"`
}

type WorkloadBytes struct {
	Read    *WorkloadMetrics `json:"read"`
	Written *WorkloadMetrics `json:"written"`
}

type WorkloadKeys struct {
	Read *WorkloadMetrics `json:"read"`
}

type WorkloadOperations struct {
	LocationRequests *WorkloadMetrics `json:"location_requests"`
	LowPriorityReads *WorkloadMetrics `json:"low_priority_reads"`
	MemoryErrors     *WorkloadMetrics `json:"memory_errors"`
	ReadRequests     *WorkloadMetrics `json:"read_requests"`
	Reads            *WorkloadMetrics `json:"reads"`
	Writes           *WorkloadMetrics `json:"writes"`
}

type WorkloadTransactions struct {
	Committed                *WorkloadMetrics `json:"committed"`
	Conflicted               *WorkloadMetrics `json:"conflicted"`
	RejectedForQueuedTooLong *WorkloadMetrics `json:"rejected_for_queued_too_long"`
	Started                  *WorkloadMetrics `json:"started"`
	StartedBatchPriority     *WorkloadMetrics `json:"started_batch_priority"`
	StartedDefaultPriority   *WorkloadMetrics `json:"started_default_priority"`
	StartedImmediatePriority *WorkloadMetrics `json:"started_immediate_priority"`
}
