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
	assert.Equal(t, status.Cluster.Data.AveragePartitionSizeBytes, 20668340)
	assert.Equal(t, status.Cluster.Data.LeastOperatingSpaceBytesLogServer, 604285174979)
	assert.Equal(t, status.Cluster.Data.LeastOperatingSpaceBytesStorageServer, 854875383)
	assert.Equal(t, status.Cluster.Data.PartitionsCount, 2)
	assert.Equal(t, status.Cluster.Data.TotalDiskUsedBytes, 542511104)
	assert.Equal(t, status.Cluster.Data.TotalKvSizeBytes, 57120499)
}

func TestMovingDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, status.Cluster.Data.MovingData.HighestPriority, 0)
	assert.Equal(t, status.Cluster.Data.MovingData.InFlightBytes, 0)
	assert.Equal(t, status.Cluster.Data.MovingData.InQueueBytes, 0)
	assert.Equal(t, status.Cluster.Data.MovingData.TotalWrittenBytes, 0)
}

func TestStateSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, status.Cluster.Data.State.Description, "")
	assert.Equal(t, status.Cluster.Data.State.MinReplicasRemaining, 1)
	assert.Equal(t, status.Cluster.Data.State.Healthy, true)
	assert.Equal(t, status.Cluster.Data.State.Name, "healthy")
}

func TestTeamTrackersSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, len(status.Cluster.Data.TeamTrackers), 1)
	assert.Equal(t, status.Cluster.Data.TeamTrackers[0].InFlightBytes, 0)
	assert.True(t, status.Cluster.Data.TeamTrackers[0].Primary)
	assert.True(t, status.Cluster.Data.TeamTrackers[0].State.Healthy)
	assert.Equal(t, status.Cluster.Data.TeamTrackers[0].State.MinReplicasRemaining, 1)
	assert.Equal(t, status.Cluster.Data.TeamTrackers[0].State.Name, "healthy")
}
