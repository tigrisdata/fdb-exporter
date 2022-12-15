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

type Log struct {
	BeginVersion         int            `json:"begin_version"`
	Current              bool           `json:"current"`
	Epoch                int            `json:"epoch"`
	LogFaultTolerance    int            `json:"log_fault_tolerance"`
	LogInterfaces        []LogInterface `json:"log_interfaces"`
	LogReplicationFactor int            `json:"log_replication_factor"`
	LogWriteAntiQuorum   int            `json:"log_write_anti_quorum"`
	PossiblyLosingData   bool           `json:"possibly_losing_data"`
}

type LogInterface struct {
	Address string `json:"address"`
	Healthy bool   `json:"healthy"`
	Id      string `json:"id"`
}
