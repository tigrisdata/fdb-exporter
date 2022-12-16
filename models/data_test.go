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

func TestDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	data := status.Cluster.Data
	assert.Equal(t, data.AveragePartitionSizeBytes, 20668340)
	assert.Equal(t, data.LeastOperatingSpaceBytesLogServer, 604285174979)
	assert.Equal(t, data.LeastOperatingSpaceBytesStorageServer, 854875383)
	assert.Equal(t, data.PartitionsCount, 2)
	assert.Equal(t, data.TotalDiskUsedBytes, 542511104)
	assert.Equal(t, data.TotalKvSizeBytes, 57120499)
}

func TestMovingDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	movingData := status.Cluster.Data.MovingData
	assert.Equal(t, movingData.HighestPriority, 0)
	assert.Equal(t, movingData.InFlightBytes, 0)
	assert.Equal(t, movingData.InQueueBytes, 0)
	assert.Equal(t, movingData.TotalWrittenBytes, 0)
}

func TestStateSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	dataState := status.Cluster.Data.State
	assert.Equal(t, dataState.Description, "")
	assert.Equal(t, dataState.MinReplicasRemaining, 1)
	assert.Equal(t, dataState.Healthy, true)
	assert.Equal(t, dataState.Name, "healthy")
}

func TestTeamTrackersSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, len(status.Cluster.Data.TeamTrackers), 1)
	teamTracker := status.Cluster.Data.TeamTrackers[0]
	assert.Equal(t, teamTracker.InFlightBytes, 0)
	assert.True(t, teamTracker.Primary)
	assert.True(t, teamTracker.State.Healthy)
	assert.Equal(t, teamTracker.State.MinReplicasRemaining, 1)
	assert.Equal(t, teamTracker.State.Name, "healthy")
}
