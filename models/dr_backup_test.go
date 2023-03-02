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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDRBackupSource(t *testing.T) {
	status := CheckJsonFile(t, "dr-backup-source.json")
	drBackup := status.Cluster.Layers.DrBackup
	defaultTag := drBackup.Tags["default"]
	drBackupInstances := drBackup.Instances["1427712e65c40eb5d6e18d32be1613a9"]
	assert.Equal(t, defaultTag.RunningBackup, true)
	assert.Equal(t, defaultTag.RunningBackupIsRestorable, true)
	assert.Equal(t, defaultTag.BackupState, "is differential")
	assert.Equal(t, defaultTag.SecondsBehind, 0.18545900000000001)
	assert.Equal(t, drBackupInstances.Version, "7.1.7")
}

func TestDRBackupDest(t *testing.T) {
	status := CheckJsonFile(t, "dr-backup-dest.json")
	drBackup := status.Cluster.Layers.DrBackupDest
	defaultTag := drBackup.Tags["default"]
	drBackupInstances := drBackup.Instances["1427712e65c40eb5d6e18d32be1613a9"]
	assert.Equal(t, defaultTag.RunningBackup, true)
	assert.Equal(t, defaultTag.RunningBackupIsRestorable, true)
	assert.Equal(t, defaultTag.BackupState, "is differential")
	assert.Equal(t, defaultTag.SecondsBehind, 69.928753999999998)
	assert.Equal(t, drBackupInstances.Version, "7.1.7")
}
