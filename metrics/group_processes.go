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

package metrics

import (
	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/models"
)

type ProcessesMetricGroup struct {
	metricGroup
}

func NewProcessesMetricGroup(reporter *MetricReporter) *ProcessesMetricGroup {
	parentScope := reporter.GetScopeOrExit("cluster")
	p := &ProcessesMetricGroup{*newMetricGroup("processes", parentScope, reporter)}
	p.AddScope(parentScope, "grv_lat", "grv")
	p.AddScope(parentScope, "commit_lat", "commit")
	p.AddScope(parentScope, "read_lat", "read")
	return p
}

func (p *ProcessesMetricGroup) getTags(processName string, process *models.Process) map[string]string {
	tags := GetBaseTags()

	// Try to tag the name with the human readable name (as published by the operator) instead of the key in the json
	if process.Locality != nil {
		tags["fdb_pod_name"] = process.Locality.InstanceId
	} else {
		tags["fdb_pod_name"] = processName
	}
	tags["class_type"] = process.ClassType
	// On the same tally scope the same tag set should be present
	tags["storage"] = "0"
	tags["log"] = "0"
	tags["master"] = "0"
	tags["coordinator"] = "0"
	tags["commit_proxy"] = "0"
	tags["grv_proxy"] = "0"
	tags["cluster_controller"] = "0"
	tags["data_distributor"] = "0"
	tags["ratekeeper"] = "0"
	tags["resolver"] = "0"
	for _, role := range process.Roles {
		switch role.Role {
		case "storage":
			tags["storage"] = "1"
		case "log":
			tags["log"] = "1"
		case "master":
			tags["master"] = "1"
		case "coordinator":
			tags["coordinator"] = "1"
		case "commit_proxy":
			tags["commit_proxy"] = "1"
		case "grv_proxy":
			tags["grv_proxy"] = "1"
		case "cluster_controller":
			tags["cluster_controller"] = "1"
		case "data_distributor":
			tags["data_distributor"] = "1"
		case "ratekeeper":
			tags["ratekeeper"] = "1"
		case "resolver":
			tags["resolver"] = "1"
		}
	}
	return tags
}

func (p *ProcessesMetricGroup) getLatencyTags(processName string, process *models.Process, quantile string) map[string]string {
	tags := p.getTags(processName, process)
	tags["quantile"] = quantile
	return tags
}

func (p *ProcessesMetricGroup) getLatencyTagsWithPriority(processName string, process *models.Process, quantile string, priority string) map[string]string {
	tags := p.getLatencyTags(processName, process, quantile)
	tags["priority"] = priority
	return tags
}

func (p *ProcessesMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := p.GetScopeOrExit("default")
	if !isValidProcesses(status) {
		log.Error().Msg("failed to get processes metric group")
		return
	}
	processes := status.Cluster.Processes
	for processName, process := range processes {
		metrics := make(map[string]interface{})
		tags := p.getTags(processName, &process)
		metrics["excluded"] = process.Excluded
		if process.Cpu != nil {
			metrics["cpu_cores"] = process.Cpu.UsageCores
		}
		if process.Disk != nil {
			metrics["disk_busy"] = process.Disk.Busy
			metrics["disk_free"] = process.Disk.FreeBytes
			metrics["disk_total_bytes"] = process.Disk.TotalBytes
			if process.Disk.Reads != nil {
				metrics["disk_reads_count"] = process.Disk.Reads.Counter
				metrics["disk_reads_hz"] = process.Disk.Reads.Hz
			}
			if process.Disk.Writes != nil {
				metrics["disk_writes_count"] = process.Disk.Writes.Counter
				metrics["disk_writes_hz"] = process.Disk.Writes.Hz
			}
		}
		if process.Memory != nil {
			metrics["mem_available_bytes"] = process.Memory.AvailableBytes
			metrics["mem_limit_bytes"] = process.Memory.LimitBytes
			metrics["mem_unused_allocated_memory"] = process.Memory.UnusedAllocatedMemory
			metrics["mem_unused_bytes"] = process.Memory.UsedBytes
		}
		if process.Network != nil {
			metrics["network_conn_errors_hz"] = process.Network.ConnectionErrors.Hz
			metrics["network_conn_closed_hz"] = process.Network.ConnectionsClosed.Hz
			metrics["network_conn_established"] = process.Network.ConnectionsEstablished.Hz
			metrics["network_current_connections"] = process.Network.CurrentConnections
			metrics["network_megabits_sent"] = process.Network.MegabitsSent.Hz
			metrics["network_megabits_received"] = process.Network.MegabitsReceived.Hz
		}
		for _, role := range process.Roles {
			switch role.Role {
			case "grv_proxy":
				if role.GrvLatencyStatistics != nil {
					grvLatScope := p.GetScopeOrExit("grv_lat")
					defaultMedianTags := p.getLatencyTagsWithPriority(processName, &process, "0.5", "default")
					batchMedianTags := p.getLatencyTagsWithPriority(processName, &process, "0.5", "batch")
					SetGauge(grvLatScope, "latency", defaultMedianTags, role.GrvLatencyStatistics.Default.Median)
					SetGauge(grvLatScope, "latency", batchMedianTags, role.GrvLatencyStatistics.Batch.Median)
					defaultP95Tags := p.getLatencyTagsWithPriority(processName, &process, "0.95", "default")
					batchP95Tags := p.getLatencyTagsWithPriority(processName, &process, "0.95", "batch")
					SetGauge(grvLatScope, "latency", defaultP95Tags, role.GrvLatencyStatistics.Default.P95)
					SetGauge(grvLatScope, "latency", batchP95Tags, role.GrvLatencyStatistics.Batch.P95)
					defaultP99Tags := p.getLatencyTagsWithPriority(processName, &process, "0.99", "default")
					batchP99Tags := p.getLatencyTagsWithPriority(processName, &process, "0.99", "batch")
					SetGauge(grvLatScope, "latency", defaultP99Tags, role.GrvLatencyStatistics.Default.P99)
					SetGauge(grvLatScope, "latency", batchP99Tags, role.GrvLatencyStatistics.Batch.P99)
				}
			case "commit_proxy":
				if role.CommitLatencyStatistics != nil {
					commitLatScope := p.GetScopeOrExit("commit_lat")
					medianTags := p.getLatencyTags(processName, &process, "0.5")
					SetGauge(commitLatScope, "latency", medianTags, role.CommitLatencyStatistics.Median)
					p95Tags := p.getLatencyTags(processName, &process, "0.95")
					SetGauge(commitLatScope, "latency", p95Tags, role.CommitLatencyStatistics.P95)
					p99Tags := p.getLatencyTags(processName, &process, "0.99")
					SetGauge(commitLatScope, "latency", p99Tags, role.CommitLatencyStatistics.P99)
				}
			case "storage":
				metrics["kvstore_available_bytes"] = role.KvStoreAvailableBytes
				metrics["kvstore_free_bytes"] = role.KvStoreFreeBytes
				metrics["kvstore_total_bytes"] = role.KvStoreTotalBytes
				metrics["kvstore_used_bytes"] = role.KvStoreUsedBytes
				metrics["queue_disk_available_bytes"] = role.QueueDiskAvailableBytes
				metrics["queue_disk_free_bytes"] = role.QueueDiskFreeBytes
				metrics["queue_disk_total_bytes"] = role.QueueDiskTotalBytes
				metrics["queue_disk_used_bytes"] = role.QueueDiskUsedBytes
				metrics["stored_bytes"] = role.StoredBytes
				if role.DataLag != nil {
					metrics["data_lag_seconds"] = role.DataLag.Seconds
					metrics["data_lag_versions"] = role.DataLag.Versions
				}
				if role.DurabilityLag != nil {
					metrics["durability_lag_seconds"] = role.DurabilityLag.Seconds
					metrics["durability_lag_versions"] = role.DurabilityLag.Versions
				}
				if role.TotalQueries != nil {
					metrics["query_count"] = role.TotalQueries.Counter
					metrics["query_hz"] = role.TotalQueries.Hz
				}
				if role.InputBytes != nil {
					metrics["input_bytes"] = role.InputBytes.Counter
				}
				if role.DurableBytes != nil {
					metrics["durability_bytes"] = role.DurableBytes.Counter
				}

				if role.ReadLatencyStatistics != nil {
					readLatScope := p.GetScopeOrExit("read_lat")
					medianTags := p.getLatencyTags(processName, &process, "0.5")
					SetGauge(readLatScope, "latency", medianTags, role.ReadLatencyStatistics.Median)
					p95Tags := p.getLatencyTags(processName, &process, "0.95")
					SetGauge(readLatScope, "latency", p95Tags, role.ReadLatencyStatistics.P95)
					p99Tags := p.getLatencyTags(processName, &process, "0.99")
					SetGauge(readLatScope, "latency", p99Tags, role.ReadLatencyStatistics.P99)
				}
			case "log":
				metrics["kvstore_available_bytes"] = role.KvStoreAvailableBytes
				metrics["kvstore_free_bytes"] = role.KvStoreFreeBytes
				metrics["kvstore_total_bytes"] = role.KvStoreTotalBytes
				metrics["kvstore_used_bytes"] = role.KvStoreUsedBytes
				metrics["queue_disk_available_bytes"] = role.QueueDiskAvailableBytes
				metrics["queue_disk_available_bytes"] = role.QueueDiskFreeBytes
				metrics["queue_disk_total_bytes"] = role.QueueDiskTotalBytes
				metrics["queue_disk_used_bytes"] = role.QueueDiskUsedBytes
				if role.InputBytes != nil {
					metrics["input_bytes"] = role.InputBytes.Counter
				}
				if role.DurableBytes != nil {
					metrics["durability_bytes"] = role.DurableBytes.Counter
				}
				if role.InputBytes != nil && role.DurableBytes != nil {
					// The input bytes shows how many bytes did the log receive, the number of durable bytes
					// are already written to the storage server, if this is growing, storage servers are not taking
					// writes or not taking writes fast enough
					metrics["log_queue_length"] = role.InputBytes.Counter - role.DurableBytes.Counter
				}
			}
		}
		SetMultipleGauges(scope, metrics, tags)
	}
}
