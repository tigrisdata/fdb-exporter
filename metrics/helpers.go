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

import (
	"github.com/rs/zerolog/log"
	"github.com/uber-go/tally"
)

func SetGauge(scope tally.Scope, name string, tags map[string]string, value interface{}) {
	switch v := value.(type) {
	case bool:
		scope.Tagged(tags).Gauge(name).Update(convertBool(v))
	case int64:
		scope.Tagged(tags).Gauge(name).Update(float64(v))
	case int:
		scope.Tagged(tags).Gauge(name).Update(float64(v))
	case float64:
		scope.Tagged(tags).Gauge(name).Update(v)
	default:
		log.Error().Str("name", name).Msg("could not determine type for gauge")
	}
}

func SetMultipleGauges(scope tally.Scope, metrics map[string]interface{}, tags map[string]string) {
	for name, value := range metrics {
		SetGauge(scope, name, tags, value)
	}
}
