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
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

type WorkloadBytesMetricGroup struct {
	MetricGroup
}

func NewWorkloadBytesMetricGroup() *WorkloadBytesMetricGroup {
	w := WorkloadBytesMetricGroup{}
	w.scopes = make(map[string]tally.Scope)
	return &w
}

func (w *WorkloadBytesMetricGroup) SetStatus(status *models.FullStatus) {
	w.status = status
}

func (w *WorkloadBytesMetricGroup) InitScopes() {
	w.scopes["default"] = WorkloadScope.SubScope("keys")
}

func (w *WorkloadBytesMetricGroup) GetMetrics() {
	SetIntGauge(w.scopes["default"], "read", GetBaseTags(), w.status.Cluster.Workload.Bytes.Read.Counter)
	SetIntGauge(w.scopes["default"], "written", GetBaseTags(), w.status.Cluster.Workload.Bytes.Written.Counter)
}
