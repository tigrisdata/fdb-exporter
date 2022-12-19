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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMachineSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	machine := status.Cluster.Machines["2f768c26fc0f01ce8af5402f7779b62c"]
	assert.Equal(t, machine.Address, "127.0.0.1")
	assert.Equal(t, machine.ContributingWorkers, 1)
	assert.False(t, machine.Excluded)
}

func TestMachineCpuSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	cpu := status.Cluster.Machines["2f768c26fc0f01ce8af5402f7779b62c"].Cpu
	assert.Equal(t, cpu.LogicalCoreUtilization, 0.042182900000000002)
}

func TestMachineLocalitySingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	locality := status.Cluster.Machines["2f768c26fc0f01ce8af5402f7779b62c"].Locality
	assert.Equal(t, locality.MachineId, "2f768c26fc0f01ce8af5402f7779b62c")
	assert.Equal(t, locality.ProcessId, "dbfd37cad094516ba1ee62c6345b3469")
	assert.Equal(t, locality.ZoneId, "2f768c26fc0f01ce8af5402f7779b62c")
}

func TestMachineMemorySingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	memory := status.Cluster.Machines["2f768c26fc0f01ce8af5402f7779b62c"].Memory
	assert.Equal(t, memory.CommittedBytes, 61098893312)
	assert.Equal(t, memory.FreeBytes, 6370414592)
	assert.Equal(t, memory.TotalBytes, 67469307904)
}

func TestMachineNetworkSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	network := status.Cluster.Machines["2f768c26fc0f01ce8af5402f7779b62c"].Network
	assert.Equal(t, network.MegabitsReceived.Hz, 0.116317)
	assert.Equal(t, network.MegabitsSent.Hz, 0.116317)
	assert.Equal(t, network.TcpSegmentsRetransmitted.Hz, 0.0)
}
