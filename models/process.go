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
	Cpu         ProcessCpu       `json:"cpu"`
	Disk        ProcessDisk      `json:"disk"`
	Excluded    bool             `json:"excluded"`
	FaultDomain string           `json:"fault_domain"`
	Locality    ProcessLocality  `json:"locality"`
	MachineId   string           `json:"machine_id"`
	Memory      ProcessMemory    `json:"memory"`
	Messages    []ProcessMessage `json:"messages"`
	Network     ProcessNetwork   `json:"network"`
	Roles       []ProcessRole    `json:"roles"`
}

type ProcessCpu struct {
	UsageCores float64 `json:"usage_cores"`
}

type ProcessDisk struct {
	Busy       float64           `json:"busy"`
	FreeBytes  int               `json:"free_bytes"`
	Reads      ProcessDiskReads  `json:"reads"`
	TotalBytes int               `json:"total_bytes"`
	Writes     ProcessDiskWrites `json:"writes"`
}

type ProcessDiskReads struct {
	Counter int     `json:"counter"`
	Hz      float64 `json:"hz"`
	Sectors int     `json:"sectors"`
}

type ProcessDiskWrites struct {
	Counter int     `json:"counter"`
	Hz      float64 `json:"hz"`
	Sectors int     `json:"sectors"`
}

type ProcessLocality struct {
	InstanceId string `json:"instance_id"`
	MachineId  string `json:"machine_id"`
	ProcessId  string `json:"process_id"`
	ZoneId     string `json:"zone_id"`
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

type ProcessNetwork struct {
	ConnectionErrors       Hz  `json:"connection_errors"`
	ConnectionsClosed      Hz  `json:"connections_closed"`
	ConnectionsEstablished Hz  `json:"connections_established"`
	CurrentConnections     int `json:"current_connections"`
	MegabitsReceived       Hz  `json:"megabits_received"`
	MegabitsSent           Hz  `json:"megabits_sent"`
	TlsPolicyFailures      Hz  `json:"tls_policy_failures"`
}

type ProcessRole struct {
	// TODO: implement the process specific fields of the roles (there are lots of them)
	Id   string `json:"id"`
	Role string `json:"role"`
}
