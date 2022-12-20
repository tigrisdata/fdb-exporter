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

// Metrics for a single subsection of status json output
// Metrics in a group are collected and processed at the same time, usually they are correlated
type metricGroup struct {
	name        string
	parentScope tally.Scope
	// Pointer to the singleton MetricReporter that is created in main
	info *MetricReporter
	scoped
}

func newMetricGroup(name string, parentScope tally.Scope, mInfo *MetricReporter) *metricGroup {
	m := metricGroup{name: name, parentScope: parentScope, info: mInfo}
	m.scopes = make(map[string]tally.Scope)
	// Default scope for a metric group
	m.AddScope(parentScope, "default", name)
	return &m
}

// Each group that collects metrics needs to implement the Collectable interface
type Collectable interface {
	GetMetrics(status *models.FullStatus)
}
