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

type DbStatusMetricGroup struct {
	MetricGroup
}

func NewDbStatusMetricGroup() *DbStatusMetricGroup {
	d := DbStatusMetricGroup{}
	d.scopes = make(map[string]tally.Scope)
	return &d
}

func (d *DbStatusMetricGroup) SetStatus(status *models.FullStatus) {
	d.status = status
}

func (d *DbStatusMetricGroup) InitScopes() {
	d.scopes["default"] = ClientScope.SubScope("status")
}

func (d *DbStatusMetricGroup) GetMetrics() {
	SetBoolGauge(d.scopes["default"], "available", GetBaseTags(), d.status.Client.DatabaseStatus.Available)
	SetBoolGauge(d.scopes["default"], "healthy", GetBaseTags(), d.status.Client.DatabaseStatus.Healthy)
}
