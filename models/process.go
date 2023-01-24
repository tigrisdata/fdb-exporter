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

type Process struct {
	Address     string           `json:"address"`
	ClassSource string           `json:"class_source"`
	ClassType   string           `json:"class_type"`
	CommandLine string           `json:"command_line"`
	Cpu         *ProcessCpu      `json:"cpu"`
	Disk        *ProcessDisk     `json:"disk"`
	Excluded    bool             `json:"excluded"`
	FaultDomain string           `json:"fault_domain"`
	Locality    *ProcessLocality `json:"locality"`
	MachineId   string           `json:"machine_id"`
	Memory      *ProcessMemory   `json:"memory"`
	Messages    []ProcessMessage `json:"messages"`
	Network     *ProcessNetwork  `json:"network"`
	Roles       []ProcessRole    `json:"roles"`
}

type ProcessCpu struct {
	UsageCores float64 `json:"usage_cores"`
}

type ProcessDisk struct {
	Busy       float64             `json:"busy"`
	FreeBytes  int                 `json:"free_bytes"`
	Reads      *ProcessDiskCounter `json:"reads"`
	TotalBytes int                 `json:"total_bytes"`
	Writes     *ProcessDiskCounter `json:"writes"`
}

type ProcessDiskCounter struct {
	Counter int     `json:"counter"`
	Hz      float64 `json:"hz"`
	Sectors float64 `json:"sectors"`
}

type ProcessLocality struct {
	InstanceId string `json:"instance_id"`
	MachineId  string `json:"machineid"`
	ProcessId  string `json:"processid"`
	ZoneId     string `json:"zoneid"`
}

type ProcessMemory struct {
	AvailableBytes        int `json:"available_bytes"`
	LimitBytes            int `json:"limit_bytes"`
	UnusedAllocatedMemory int `json:"unused_allocated_memory"`
	UsedBytes             int `json:"used_bytes"`
}

type ProcessMessage struct {
	Description   string `json:"description"`
	Name          string `json:"name"`
	RawLogMessage string `json:"raw_log_message"`
	Time          int    `json:"time"`
	Type          string `json:"type"`
}

type LatencyStats struct {
	Count  int     `json:"count"`
	Max    float64 `json:"max"`
	Mean   float64 `json:"mean"`
	Median float64 `json:"median"`
	Min    float64 `json:"min"`
	P25    float64 `json:"p25"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p99.9"`
}

type ProcessNetwork struct {
	ConnectionErrors       *Hz `json:"connection_errors"`
	ConnectionsClosed      *Hz `json:"connections_closed"`
	ConnectionsEstablished *Hz `json:"connections_established"`
	CurrentConnections     int `json:"current_connections"`
	MegabitsReceived       *Hz `json:"megabits_received"`
	MegabitsSent           *Hz `json:"megabits_sent"`
	TlsPolicyFailures      *Hz `json:"tls_policy_failures"`
}

type GrvLatencyStats struct {
	Batch   *LatencyStats `json:"batch"`
	Default *LatencyStats `json:"default"`
}

type ProcessRole struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	// GRV proxy specific
	GrvLatencyStatistics *GrvLatencyStats `json:"grv_latency_statistics"`
	// Commit Proxy specific
	CommitLatencyStatistics  *LatencyStats `json:"commit_latency_statistics"`
	CommitBatchingWindowSize *LatencyStats `json:"commit_batching_window_size"`
	// Storage and Log specific
	KvStoreAvailableBytes   int64            `json:"kvstore_available_bytes"`
	KvStoreFreeBytes        int64            `json:"kvstore_free_bytes"`
	KvStoreTotalBytes       int64            `json:"kvstore_total_bytes"`
	KvStoreUsedBytes        int64            `json:"kvstore_used_bytes"`
	QueueDiskAvailableBytes int64            `json:"queue_disk_available_bytes"`
	QueueDiskFreeBytes      int64            `json:"queue_disk_free_bytes"`
	QueueDiskTotalBytes     int64            `json:"queue_disk_total_bytes"`
	QueueDiskUsedBytes      int64            `json:"queue_disk_used_bytes"`
	DataVersion             int64            `json:"data_version"`
	DurableBytes            *WorkloadMetrics `json:"durable_bytes"`
	InputBytes              *WorkloadMetrics `json:"input_bytes"`
	// Only storage specific
	BytesQueried          *WorkloadMetrics `json:"bytes_queried"`
	DataLag               *Lag             `json:"data_lag"`
	DurabilityLag         *Lag             `json:"durability_lag"`
	DurableVersion        int64            `json:"durable_version"`
	FetchedVersions       *WorkloadMetrics `json:"fetched_versions"`
	FinishedQueries       *WorkloadMetrics `json:"finished_queries"`
	KeysQueried           *WorkloadMetrics `json:"keys_queried"`
	LocalRate             int64            `json:"local_rate"`
	LowPriorityQueries    *WorkloadMetrics `json:"low_priority_queries"`
	MutationBytes         *WorkloadMetrics `json:"mutation_bytes"`
	Mutations             *WorkloadMetrics `json:"mutations"`
	QueryQueueMax         int64            `json:"query_queue_max"`
	ReadLatencyStatistics *LatencyStats    `json:"read_latency_statistics"`
	StoredBytes           int64            `json:"stored_bytes"`
	TotalQueries          *WorkloadMetrics `json:"total_queries"`
}
