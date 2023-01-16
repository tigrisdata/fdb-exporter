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

package models

type Qos struct {
	BatchPerformanceLimitedBy          PerformanceLimitedBy `json:"batch_performance_limited_by"`
	BatchReleasedTransactionsPerSecond float64              `json:"batch_released_transactions_per_second"`
	BatchTransactionsPerSecondLimit    float64              `json:"batch_transactions_per_second_limit"`
	LimitingDataLagStorageServer       Lag                  `json:"limiting_data_lag_storage_server"`
	LimitingDurabilityLagStorageServer Lag                  `json:"limiting_durability_lag_storage_server"`
	LimitingQueueBytesStorageServer    int64                `json:"limiting_queue_bytes_storage_server"`
	PerformanceLimitedBy               PerformanceLimitedBy `json:"performance_limited_by"`
	ReleasedTransactionsPerSecond      float64              `json:"released_transactions_per_second"`
	ThrottledTags                      ThrottledTags        `json:"throttled_tags"`
	TransactionsPerSecondLimit         float64              `json:"transactions_per_second_limit"`
	WorstDataLagStorageServer          Lag                  `json:"worst_data_lag_storage_server"`
	WorstDurabilityLagStorageServer    Lag                  `json:"worst_durability_lag_storage_server"`
	WorstQueueBytesLogServer           int64                `json:"worst_queue_bytes_log_server"`
	WorstQueueBytesStorageServer       int64                `json:"worst_queue_bytes_storage_server"`
}

type PerformanceLimitedBy struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	ReasonId    int64  `json:"reason_id"`
}

type ThrottledTags struct {
	Auto   AutoThrottledTags   `json:"auto"`
	Manual ManualThrottledTags `json:"manual"`
}

type AutoThrottledTags struct {
	BusyRead        int64 `json:"busy_read"`
	BusyWrite       int64 `json:"busy_write"`
	Count           int64 `json:"count"`
	RecommendedOnly int64 `json:"recommended_only"`
}

type ManualThrottledTags struct {
	Count int64 `json:"count"`
}
