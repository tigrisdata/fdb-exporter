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

package metrics

import (
	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/models"
)

type BackupMetricGroup struct {
	metricGroup
}

func NewBackupMetricGroup(reporter *MetricReporter) *BackupMetricGroup {
	parentScope := reporter.GetScopeOrExit("cluster")
	b := &BackupMetricGroup{*newMetricGroup("backup", parentScope, reporter)}
	// In the future, other sections from the backup key might be processed here
	b.AddScope(parentScope, "backup_tag", "backup_tag")
	b.AddScope(parentScope, "backup_instances", "backup_instances")
	b.AddScope(parentScope, "backup_config", "backup_config")
	return b
}

func (b *BackupMetricGroup) GetMetrics(status *models.FullStatus) {
	b.getNoBackupMetrics(status)
	b.getTaggedMetrics(status)
	b.getInstanceMetrics(status)
}

func (b *BackupMetricGroup) emitNoBackupMetrics() {
	taggedScope := b.GetScopeOrExit("backup_config")
	SetGauge(taggedScope, "absent", GetBaseTags(), 1)
}

func (b *BackupMetricGroup) getNoBackupMetrics(status *models.FullStatus) {
	// If the backup section is not present emit fdb_cluster_backup_config_absent metric
	if !isValidBackup(status) {
		// Emit if there is no backup section at all
		b.emitNoBackupMetrics()
	} else {
		if status.Cluster.Layers.Backup.Tags == nil {
			// Emit also if there is no tag section. That can happen when backup agents are connected, but nothing
			// is configured.
			b.emitNoBackupMetrics()
		}
	}
}

func (b *BackupMetricGroup) getTaggedMetrics(status *models.FullStatus) {
	if !isValidBackup(status) {
		log.Error().Msg("Failed to get backup tag metric group")
		return
	}
	backup := status.Cluster.Layers.Backup
	taggedScope := b.GetScopeOrExit("backup_tag")
	tagMetrics := make(map[string]interface{})
	taggedBackupTags := GetBaseTags()
	for tagName, backupTag := range backup.Tags {
		taggedBackupTags["backup_tag"] = tagName
		tagMetrics["is_running"] = backupTag.RunningBackup
		tagMetrics["running_is_restorable"] = backupTag.RunningBackupIsRestorable
		tagMetrics["last_restorable_seconds_behind"] = backupTag.LastRestorableSecondsBehind
		SetMultipleGauges(taggedScope, tagMetrics, taggedBackupTags)
	}
}

func (b *BackupMetricGroup) getInstanceMetrics(status *models.FullStatus) {
	if !isValidBackup(status) {
		log.Error().Msg("Failed to get backup instance metric group")
		return
	}
	backup := status.Cluster.Layers.Backup

	if backup.Instances != nil {
		numInstances := len(status.Cluster.Layers.Backup.Instances)
		instanceScope := b.GetScopeOrExit("backup_instances")
		SetGauge(instanceScope, "count", GetBaseTags(), numInstances)
	}
}
