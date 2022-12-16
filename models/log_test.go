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

func TestLogSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	clusterLog := status.Cluster.Logs[0]
	assert.Equal(t, clusterLog.BeginVersion, 2264599536058)
	assert.True(t, clusterLog.Current)
	assert.Equal(t, clusterLog.Epoch, 58)
	assert.Equal(t, clusterLog.LogFaultTolerance, 0)
	assert.Equal(t, clusterLog.LogReplicationFactor, 1)
	assert.Equal(t, clusterLog.PossiblyLosingData, false)
}