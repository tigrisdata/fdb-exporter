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
	"github.com/uber-go/tally"
)

func SetBoolGauge(scope tally.Scope, name string, tags map[string]string, value bool) {
	scope.Tagged(tags).Gauge(name).Update(convertBool(value))
}

func SetIntGauge(scope tally.Scope, name string, tags map[string]string, value int) {
	scope.Tagged(tags).Gauge(name).Update(float64(value))
}

func SetFloatGauge(scope tally.Scope, name string, tags map[string]string, value float64) {
	scope.Tagged(tags).Gauge(name).Update(value)
}
