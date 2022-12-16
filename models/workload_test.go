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

func TestWorkloadBytesSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	workloadBytes := status.Cluster.Workload.Bytes
	assert.Equal(t, workloadBytes.Read.Counter, 1521265096)
	assert.Equal(t, workloadBytes.Read.Hz, 3361.9899999999998)
	assert.Equal(t, workloadBytes.Read.Roughness, 1333.3800000000001)
	assert.Equal(t, workloadBytes.Written.Counter, 15203569)
	assert.Equal(t, workloadBytes.Written.Hz, 0.0)
	assert.Equal(t, workloadBytes.Written.Roughness, 0.0)
}

func TestWorkloadKeysSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	workloadKeys := status.Cluster.Workload.Keys
	assert.Equal(t, workloadKeys.Read.Counter, 4192507)
	assert.Equal(t, workloadKeys.Read.Hz, 11.4)
	assert.Equal(t, workloadKeys.Read.Roughness, 3.5246599999999999)
}

func TestWorkloadOperationsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	workloadOperations := status.Cluster.Workload.Operations
	assert.Equal(t, workloadOperations.LocationRequests.Counter, 1564)
	assert.Equal(t, workloadOperations.LocationRequests.Hz, 0.79995699999999992)
	assert.Equal(t, workloadOperations.LocationRequests.Roughness, 2.9997800000000003)
	assert.Equal(t, workloadOperations.LowPriorityReads.Counter, 0)
	assert.Equal(t, workloadOperations.LowPriorityReads.Hz, 0.0)
	assert.Equal(t, workloadOperations.LowPriorityReads.Roughness, 0.0)
	assert.Equal(t, workloadOperations.MemoryErrors.Counter, 0)
	assert.Equal(t, workloadOperations.MemoryErrors.Hz, 0.0)
	assert.Equal(t, workloadOperations.MemoryErrors.Roughness, 0.0)
	assert.Equal(t, workloadOperations.ReadRequests.Counter, 8821994)
	assert.Equal(t, workloadOperations.ReadRequests.Hz, 21.1999)
	assert.Equal(t, workloadOperations.ReadRequests.Roughness, 7.0466199999999999)
	assert.Equal(t, workloadOperations.Reads.Counter, 8821994)
	assert.Equal(t, workloadOperations.Reads.Hz, 21.1999)
	assert.Equal(t, workloadOperations.Reads.Roughness, 7.0472700000000001)
	assert.Equal(t, workloadOperations.Writes.Counter, 114099)
	assert.Equal(t, workloadOperations.Writes.Hz, 0.0)
	assert.Equal(t, workloadOperations.Writes.Roughness, 0.0)
}

func TestWorkloadTransactionsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	workloadTransactions := status.Cluster.Workload.Transactions
	assert.Equal(t, workloadTransactions.Committed.Counter, 64050)
	assert.Equal(t, workloadTransactions.Committed.Hz, 0.199989)
	assert.Equal(t, workloadTransactions.Committed.Roughness, 0.0)
	assert.Equal(t, workloadTransactions.Conflicted.Counter, 22)
	assert.Equal(t, workloadTransactions.Conflicted.Hz, 0.0)
	assert.Equal(t, workloadTransactions.Conflicted.Roughness, 0.0)
	assert.Equal(t, workloadTransactions.RejectedForQueuedTooLong.Counter, 0)
	assert.Equal(t, workloadTransactions.RejectedForQueuedTooLong.Hz, 0.0)
	assert.Equal(t, workloadTransactions.RejectedForQueuedTooLong.Roughness, 0.0)
	assert.Equal(t, workloadTransactions.Started.Counter, 2473210)
	assert.Equal(t, workloadTransactions.Started.Hz, 8.9984500000000001)
	assert.Equal(t, workloadTransactions.Started.Roughness, 2.2996500000000002)
	assert.Equal(t, workloadTransactions.StartedBatchPriority.Counter, 637)
	assert.Equal(t, workloadTransactions.StartedBatchPriority.Hz, 0.199966)
	assert.Equal(t, workloadTransactions.StartedBatchPriority.Roughness, 0.0)
	assert.Equal(t, workloadTransactions.StartedDefaultPriority.Counter, 1909130)
	assert.Equal(t, workloadTransactions.StartedDefaultPriority.Hz, 5.9989699999999999)
	assert.Equal(t, workloadTransactions.StartedDefaultPriority.Roughness, 1.20198)
	assert.Equal(t, workloadTransactions.StartedImmediatePriority.Counter, 563443)
	assert.Equal(t, workloadTransactions.StartedImmediatePriority.Hz, 2.7995200000000002)
	assert.Equal(t, workloadTransactions.StartedImmediatePriority.Roughness, 1.5168699999999999)
}
