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

func TestBackupRunning(t *testing.T) {
	status := CheckJsonFile(t, "backup-running.json")
	backup := status.Cluster.Layers.Backup
	defaultTag := backup.Tags["default"]
	backupInstances := backup.Instances["a7c44a12377abe7a35362b8a520eac5a"]
	assert.Equal(t, defaultTag.RunningBackup, true)
	assert.Equal(t, defaultTag.RunningBackupIsRestorable, false)
	assert.Equal(t, defaultTag.LastRestorableSecondsBehind, 54539.248951000001)
	assert.Equal(t, defaultTag.LastRestorableVersion, int64(18979857871))
	assert.Equal(t, defaultTag.CurrentStats, "has been started")
	assert.Equal(t, backupInstances.Version, "7.1.7")
}

func TestBackupCompleted(t *testing.T) {
	status := CheckJsonFile(t, "backup-completed.json")
	backup := status.Cluster.Layers.Backup
	defaultTag := status.Cluster.Layers.Backup.Tags["default"]
	backupInstances := backup.Instances["a7c44a12377abe7a35362b8a520eac5a"]
	assert.Equal(t, defaultTag.RunningBackup, false)
	assert.Equal(t, defaultTag.RunningBackupIsRestorable, false)
	assert.Equal(t, defaultTag.LastRestorableSecondsBehind, 46.779859000000002)
	assert.Equal(t, defaultTag.LastRestorableVersion, int64(73551594910))
	assert.Equal(t, defaultTag.CurrentStats, "has been completed")
	assert.Equal(t, backupInstances.Version, "7.1.7")
}
