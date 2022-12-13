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

import "github.com/tigrisdata/fdb-exporter/models"

type DbStatusMetricGroup struct {
	MetricGroup
}

func NewDbStatusMetricGroup(mInfo *MetricInfo) *DbStatusMetricGroup {
	return &DbStatusMetricGroup{*NewMetricGroup("status", mInfo.scopes["client"], mInfo)}
}

func (d *DbStatusMetricGroup) GetMetrics(status *models.FullStatus) {
	SetBoolGauge(d.scopes["default"], "available", GetBaseTags(), status.Client.DatabaseStatus.Available)
	SetBoolGauge(d.scopes["default"], "healthy", GetBaseTags(), status.Client.DatabaseStatus.Healthy)
}
