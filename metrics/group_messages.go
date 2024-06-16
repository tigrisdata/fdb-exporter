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

package metrics

import "github.com/tigrisdata/fdb-exporter/models"

type ClusterMessageMetricGroup struct {
	metricGroup
}

func NewClusterMessageMetricGroup(reporter *MetricReporter) *ClusterMessageMetricGroup {
	parentScope := reporter.GetScopeOrExit("cluster")
	p := &ClusterMessageMetricGroup{*newMetricGroup("messages", parentScope, reporter)}
	p.AddScope(parentScope, "global", "global")
	return p
}

func (c *ClusterMessageMetricGroup) getTags(name string) map[string]string {
	tags := GetBaseTags()
	tags["name"] = name
	return tags
}

func (c *ClusterMessageMetricGroup) GetMetrics(status *models.FullStatus) {
	messagesCounter := make(map[string]int)
	for _, clusterMessage := range status.Cluster.Messages {
		messagesCounter[clusterMessage.Name] += 1
	}
	for name, count := range messagesCounter {
		SetGauge(c.GetScopeOrExit("global"), "messages", c.getTags(name), count)
	}
}
