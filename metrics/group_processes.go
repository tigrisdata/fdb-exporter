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
	return &ProcessesMetricGroup{*newMetricGroup("processes", reporter.GetScopeOrExit("cluster"), reporter)}
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
		case "cluster_contoller":
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
		SetMultipleGauges(scope, metrics, tags)
	}
}
