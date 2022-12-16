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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	process := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"]
	assert.Equal(t, process.Address, "127.0.0.1:4689")
	assert.Equal(t, process.ClassSource, "command_line")
	assert.Equal(t, process.ClassType, "unset")
	assert.Equal(t, process.CommandLine, "/usr/local/libexec/fdbserver --cluster_file=/usr/local/etc/foundationdb/fdb.cluster --datadir=/usr/local/foundationdb/data/4689 --listen_address=public --logdir=/usr/local/foundationdb/logs --public_address=auto:4689")
	assert.False(t, process.Excluded)
	assert.Equal(t, process.FaultDomain, "2f768c26fc0f01ce8af5402f7779b62c")
	assert.Equal(t, process.MachineId, "2f768c26fc0f01ce8af5402f7779b62c")
}

func TestProcessCpuSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	process := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"]
	assert.Equal(t, process.Cpu.UsageCores, 0.0277308)
}

func TestProcessDiskSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processDisk := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Disk
	assert.Equal(t, processDisk.Busy, 0.022398300000000003)
	assert.Equal(t, processDisk.FreeBytes, 654297223168)
	assert.Equal(t, processDisk.TotalBytes, 1000240963584)
}

func TestProcessDiskReadsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processDiskReads := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Disk.Reads
	assert.Equal(t, processDiskReads.Counter, 19058541)
	assert.Equal(t, processDiskReads.Hz, 52.995899999999999)
	assert.Equal(t, processDiskReads.Sectors, 0)
}

func TestProcessDiskWritesSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processDiskWrites := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Disk.Writes
	assert.Equal(t, processDiskWrites.Counter, 42503287)
	assert.Equal(t, processDiskWrites.Hz, 62.795099999999998)
	assert.Equal(t, processDiskWrites.Sectors, 0)
}

func TestProcessLocalityWritesSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processLocality := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Locality
	assert.Equal(t, processLocality.MachineId, "2f768c26fc0f01ce8af5402f7779b62c")
	assert.Equal(t, processLocality.ProcessId, "dbfd37cad094516ba1ee62c6345b3469")
	assert.Equal(t, processLocality.ZoneId, "2f768c26fc0f01ce8af5402f7779b62c")
}

func TestProcessMemoryWritesSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processMemory := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Memory
	assert.Equal(t, processMemory.AvailableBytes, 8589934592)
	assert.Equal(t, processMemory.LimitBytes, 8589934592)
	assert.Equal(t, processMemory.UnusedAllocatedMemory, 262144)
	assert.Equal(t, processMemory.UsedBytes, 37489426432)
}

func TestProcessMessageSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	process := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"]
	assert.Equal(t, len(process.Messages), 0)
}

func TestProcessNetworkSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processNetwork := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Network
	assert.Equal(t, processNetwork.ConnectionErrors.Hz, 0.0)
	assert.Equal(t, processNetwork.ConnectionsClosed.Hz, 0.0)
	assert.Equal(t, processNetwork.ConnectionsEstablished.Hz, 0.199985)
	assert.Equal(t, processNetwork.CurrentConnections, 2)
	assert.Equal(t, processNetwork.MegabitsReceived.Hz, 0.043350300000000001)
	assert.Equal(t, processNetwork.MegabitsSent.Hz, 0.054696600000000005)
	assert.Equal(t, processNetwork.TlsPolicyFailures.Hz, 0.0)
}

func TestProcessRoleSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	processRoles := status.Cluster.Processes["dbfd37cad094516ba1ee62c6345b3469"].Roles
	assert.Equal(t, processRoles[0].Id, "40da69ee69f3a89b")
	assert.Equal(t, processRoles[0].Role, "master")
	assert.Equal(t, processRoles[1].Id, "3526c48e754c9999")
	assert.Equal(t, processRoles[1].Role, "cluster_controller")
	assert.Equal(t, processRoles[2].Id, "64292680c8f7b098")
	assert.Equal(t, processRoles[2].Role, "data_distributor")
	assert.Equal(t, processRoles[3].Id, "daaf66c325fe5067")
	assert.Equal(t, processRoles[3].Role, "ratekeeper")
	assert.Equal(t, processRoles[4].Role, "coordinator")
}
