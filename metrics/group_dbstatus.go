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

type DbStatusMetricGroup struct {
	metricGroup
}

func NewDbStatusMetricGroup(info *MetricReporter) *DbStatusMetricGroup {
	return &DbStatusMetricGroup{*newMetricGroup("status", info.GetScopeOrExit("client"), info)}
}

func (d *DbStatusMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := d.GetScopeOrExit("default")
	if !isValidClient(status) || status.Client.DatabaseStatus == nil {
		log.Error().Msg("failed to get database status")
		return
	}
	SetGauge(scope, "available", GetBaseTags(), status.Client.DatabaseStatus.Available)
	SetGauge(scope, "healthy", GetBaseTags(), status.Client.DatabaseStatus.Healthy)
}
