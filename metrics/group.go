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

// TODO: Metric group and metric interfaces, each one will plug into these
type MetricGroup struct {
	name        string
	scopes      map[string]tally.Scope
	parentScope tally.Scope
	mInfo       *MetricInfo
}

func NewMetricGroup(name string, parentScope tally.Scope, mInfo *MetricInfo) *MetricGroup {
	m := MetricGroup{name: name, parentScope: parentScope, mInfo: mInfo}
	m.scopes = make(map[string]tally.Scope)
	m.scopes["default"] = m.parentScope.SubScope(name)
	return &m
}

func (m *MetricGroup) AddScope(parentScope tally.Scope, name string) {
	m.scopes[name] = parentScope.SubScope(name)
}

type Collectable interface {
	GetMetrics(status *models.FullStatus)
}
