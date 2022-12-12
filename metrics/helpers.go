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
