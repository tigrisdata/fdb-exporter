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

type Backup struct {
	BlobRecentIo     *BlobRecentIo             `json:"blob_recent_io"`
	Instances        map[string]BackupInstance `json:"instances"`
	InstancesRunning int64                     `json:"instances_running"`
	LastUpdated      float64                   `json:"last_updated"`
	Paused           bool                      `json:"paused"`
	Tags             map[string]BackupTag      `json:"tags"`
	TotalWorkers     int                       `json:"total_workers"`
}

type BlobRecentIo struct {
	BytesPerSecond     float64 `json:"bytes_per_second"`
	BytesSent          int64   `json:"bytes_sent"`
	RequestsFailed     int64   `json:"requests_failed"`
	RequestsSuccessful int64   `json:"requests_successful"`
}

type BackupInstance struct {
	BlobStats            *BlobStats `json:"blob_stats"`
	ConfiguredWorkers    int64      `json:"configured_workers"`
	Id                   string     `json:"id"`
	LastUpdated          float64    `json:"last_updated"`
	MainThreadCpuSeconds float64    `json:"main_thread_cpu_seconds"`
	MemoryUsage          int64      `json:"memory_usage"`
	ProcessCpuSeconds    float64    `json:"process_cpu_seconds"`
	ResidentSize         int64      `json:"resident_size"`
	Version              string     `json:"version"`
}

type BlobStats struct {
	Recent BlobStatsRecent `json:"recent"`
	Total  BlobStatsTotal  `json:"total"`
}

type BlobStatsRecent struct {
	BytesPerSecond     float64 `json:"bytes_per_second"`
	BytesSent          int64   `json:"bytes_sent"`
	RequestsFailed     int64   `json:"requests_failed"`
	RequestsSuccessful int64   `json:"requests_successful"`
}

type BlobStatsTotal struct {
	BytesSent          int64 `json:"bytes_sent"`
	RequestsFailed     int64 `json:"requests_failed"`
	RequestsSuccessful int64 `json:"requests_successful"`
}

type BackupTag struct {
	CurrentContainer            string  `json:"current_container"`
	CurrentStats                string  `json:"current_status"`
	LastRestorableSecondsBehind float64 `json:"last_restorable_seconds_behind"`
	LastRestorableVersion       int64   `json:"last_restorable_version"`
	MutationLogBytesWritten     int64   `json:"mutation_log_bytes_written"`
	MutationStreamId            string  `json:"mutation_stream_id"`
	RangeBytesWritten           int64   `json:"range_bytes_written"`
	RunningBackup               bool    `json:"running_backup"`
	RunningBackupIsRestorable   bool    `json:"running_backup_is_restorable"`
}
