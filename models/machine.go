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

type Machine struct {
	Address             string           `json:"address"`
	ContributingWorkers int              `json:"contributing_workers"`
	Cpu                 *MachineCpu      `json:"cpu"`
	Excluded            bool             `json:"excluded"`
	Locality            *MachineLocality `json:"locality"`
	Memory              *MachineMemory   `json:"memory"`
	Network             *MachineNetwork  `json:"network"`
}

type MachineCpu struct {
	LogicalCoreUtilization float64 `json:"logical_core_utilization"`
}

type MachineLocality struct {
	MachineId string `json:"machineid"`
	ProcessId string `json:"processid"`
	ZoneId    string `json:"zoneid"`
}

type MachineMemory struct {
	CommittedBytes int `json:"committed_bytes"`
	FreeBytes      int `json:"free_bytes"`
	TotalBytes     int `json:"total_bytes"`
}

type MachineNetwork struct {
	MegabitsReceived         *Hz `json:"megabits_received"`
	MegabitsSent             *Hz `json:"megabits_sent"`
	TcpSegmentsRetransmitted *Hz `json:"tcp_segments_retransmitted"`
}
