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

package clustermodel

type Data struct {
	AveragePartitionSizeBytes             int           `json:"average_partition_size_bytes"`
	LeastOperatingSpaceBytesLogServer     int           `json:"least_operating_space_bytes_log_server"`
	LeastOperatingSpaceBytesStorageServer int           `json:"least_operating_space_bytes_storage_server"`
	MovingData                            MovingData    `json:"moving_data"`
	PartitionsCount                       int           `json:"partitions_count"`
	State                                 State         `json:"state"`
	SystemKvSizeBytes                     int           `json:"system_kv_size_bytes"`
	TeamTrackers                          []TeamTracker `json:"team_trackers"`
	TotalDiskUsedBytes                    int           `json:"total_disk_used_bytes"`
	TotalKvSizeBytes                      int           `json:"total_kv_size_bytes"`
}

type MovingData struct {
	HighestPriority   int `json:"highest_priority"`
	InFlightBytes     int `json:"in_flight_bytes"`
	InQueueBytes      int `json:"in_queue_bytes"`
	TotalWrittenBytes int `json:"total_written_bytes"`
}

type State struct {
	Description          string `json:"description"`
	Healthy              bool   `json:"healthy"`
	MinReplicasRemaining int    `json:"min_replicas_remaining"`
	Name                 string `json:"name"`
}

type TeamTracker struct {
	InFlightBytes    int   `json:"in_flight_bytes"`
	Primary          bool  `json:"primary"`
	State            State `json:"state"`
	UnhealthyServers int   `json:"unhealthy_servers"`
}
