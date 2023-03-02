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

// Cluster top level status
type ClusterStatus struct {
	// TODO: add support for clients and detecting version compatibility
	// TODO: add support for incompatible connections
	ClusterControllerTimestamp int64              `json:"cluster_controller_timestamp"`
	Configuration              *Configuration     `json:"configuration"`
	ConnectionString           string             `json:"connection_string"`
	Data                       *Data              `json:"data"`
	DatabaseAvailable          bool               `json:"database_available"`
	DatabaseLockState          *LockState         `json:"database_lock_state"`
	DatacenterLag              Lag                `json:"datacenter_lag"`
	DegradedProcesses          int64              `json:"degraded_processes"`
	FaultTolerance             *FaultTolerance    `json:"fault_tolerance"`
	FullReplication            bool               `json:"full_replication"`
	Generation                 int64              `json:"generation"`
	LatencyProbe               *LatencyProbe      `json:"latency_probe"`
	Logs                       []Log              `json:"logs"`
	Machines                   map[string]Machine `json:"machines"`
	Messages                   []ClusterMessage   `json:"messages"`
	PageCache                  *PageCache         `json:"page_cache"`
	Processes                  map[string]Process `json:"processes"`
	ProtocolVersion            string             `json:"protocol_version"`
	Qos                        *Qos               `json:"qos"`
	RecoveryState              *RecoveryState     `json:"recovery_state"`
	Workload                   *Workload          `json:"workload"`
	Layers                     *Layers            `json:"layers"`
}

type LockState struct {
	Locked bool `json:"locked"`
}

type Lag struct {
	Seconds  float64 `json:"seconds"`
	Versions int64   `json:"versions"`
}

type FaultTolerance struct {
	MaxZoneFailuresWithoutLosingAvailability int64 `json:"max_zone_failures_without_losing_availability"`
	MaxZoneFailuresWithoutLosingData         int64 `json:"max_zone_failures_without_losing_data"`
}

type LatencyProbe struct {
	BatchPriorityTransactionStartSeconds     float64 `json:"batch_priority_transaction_start_seconds"`
	CommitSeconds                            float64 `json:"commit_seconds"`
	ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
	ReadSeconds                              float64 `json:"read_seconds"`
	TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
}

type Hz struct {
	Hz float64 `json:"hz"`
}

type PageCache struct {
	LogHitRate     float64 `json:"log_hit_rate"`
	StorageHitRate float64 `json:"storage_hit_rate"`
}

type RecoveryState struct {
	ActiveGenerations         int64   `json:"active_generations"`
	Description               string  `json:"description"`
	Name                      string  `json:"name"`
	SecondsSinceLastRecovered float64 `json:"seconds_since_last_recovered"`
}
