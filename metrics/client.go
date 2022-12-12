package metrics

import (
	"github.com/uber-go/tally"
)

var ClientScope tally.Scope

func InitClientMetrics() {
	ClientScope = FdbScope.SubScope("client")
}
