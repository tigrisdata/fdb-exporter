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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQosSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	qos := status.Cluster.Qos
	assert.Equal(t, qos.BatchReleasedTransactionsPerSecond, 0.0147207)
	assert.Equal(t, qos.BatchTransactionsPerSecondLimit, 294386000.0)
	assert.Equal(t, qos.LimitingQueueBytesStorageServer, 0)
	assert.Equal(t, qos.ReleasedTransactionsPerSecond, 3.85887)
	assert.Equal(t, qos.TransactionsPerSecondLimit, 137666000.0)
	assert.Equal(t, qos.WorstQueueBytesLogServer, 2)
	assert.Equal(t, qos.WorstQueueBytesStorageServer, 695)
}

func TestQosPerformanceLimitedBySingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	perfLimitedBy := status.Cluster.Qos.PerformanceLimitedBy
	assert.Equal(t, perfLimitedBy.Description, "The database is not being saturated by the workload.")
	assert.Equal(t, perfLimitedBy.Name, "workload")
	assert.Equal(t, perfLimitedBy.ReasonId, 6)
}

func TestQosAutoThrottledTagsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	autoThrottledTags := status.Cluster.Qos.ThrottledTags.Auto
	assert.Equal(t, autoThrottledTags.BusyRead, 0)
	assert.Equal(t, autoThrottledTags.BusyWrite, 0)
	assert.Equal(t, autoThrottledTags.Count, 0)
	assert.Equal(t, autoThrottledTags.RecommendedOnly, 0)
}

func TestQosManualThrottledTagsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, status.Cluster.Qos.ThrottledTags.Manual.Count, 0)
}
