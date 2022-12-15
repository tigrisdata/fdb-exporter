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
)

type CoordinatorMetricGroup struct {
	metricGroup
}

func NewCoordinatorMetricGroup(info *MetricInfo) *CoordinatorMetricGroup {
	return &CoordinatorMetricGroup{*newMetricGroup("coordinator", info.GetScopeOrExit("client"), info)}
}

func (c *CoordinatorMetricGroup) GetMetrics(status *models.FullStatus) {
	scope := c.GetScopeOrExit("default")
	SetBoolGauge(scope, "quorum", GetBaseTags(), status.Client.Coordinators.QuorumReachable)
	reachableCount := 0
	unreachableCount := 0
	for _, coordinator := range status.Client.Coordinators.Coordinators {
		if coordinator.Reachable {
			reachableCount += 1
		} else {
			unreachableCount += 1
		}
	}
	SetIntGauge(scope, "reachable", GetBaseTags(), reachableCount)
	SetIntGauge(scope, "unreachable", GetBaseTags(), unreachableCount)
}
