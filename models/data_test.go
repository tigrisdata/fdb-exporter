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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	data := status.Cluster.Data
	assert.Equal(t, data.AveragePartitionSizeBytes, int64(20668340))
	assert.Equal(t, data.LeastOperatingSpaceBytesLogServer, int64(604285174979))
	assert.Equal(t, data.LeastOperatingSpaceBytesStorageServer, int64(854875383))
	assert.Equal(t, data.PartitionsCount, int64(2))
	assert.Equal(t, data.TotalDiskUsedBytes, int64(542511104))
	assert.Equal(t, data.TotalKvSizeBytes, int64(57120499))
}

func TestMovingDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	movingData := status.Cluster.Data.MovingData
	assert.Equal(t, movingData.HighestPriority, int64(0))
	assert.Equal(t, movingData.InFlightBytes, int64(0))
	assert.Equal(t, movingData.InQueueBytes, int64(0))
	assert.Equal(t, movingData.TotalWrittenBytes, int64(0))
}

func TestStateSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	dataState := status.Cluster.Data.State
	assert.Equal(t, dataState.Description, "")
	assert.Equal(t, dataState.MinReplicasRemaining, int64(1))
	assert.Equal(t, dataState.Healthy, true)
	assert.Equal(t, dataState.Name, "healthy")
}

func TestTeamTrackersSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, len(status.Cluster.Data.TeamTrackers), 1)
	teamTracker := status.Cluster.Data.TeamTrackers[0]
	assert.Equal(t, teamTracker.InFlightBytes, int64(0))
	assert.True(t, teamTracker.Primary)
	assert.True(t, teamTracker.State.Healthy)
	assert.Equal(t, teamTracker.State.MinReplicasRemaining, int64(1))
	assert.Equal(t, teamTracker.State.Name, "healthy")
}
