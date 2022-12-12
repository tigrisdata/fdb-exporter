package metrics

import (
	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
	"io"
	"time"
)

var RootScope tally.Scope
var FdbScope tally.Scope
var Reporter promreporter.Reporter
var AllMetricGroups []Collectable

func InitMetrics() io.Closer {
	var closer io.Closer
	Reporter = promreporter.NewReporter(promreporter.Options{})
	RootScope, closer = tally.NewRootScope(tally.ScopeOptions{
		Tags:           GetBaseTags(),
		CachedReporter: Reporter,
		Separator:      promreporter.DefaultSeparator,
	}, 1*time.Second)

	// Top level tally scopes
	FdbScope = RootScope.SubScope("fdb")
	InitClientMetrics()
	InitClusterMetrics()

	AllMetricGroups = []Collectable{
		NewCoordinatorMetricGroup(),
		NewDbStatusMetricGroup(),
		NewWorkloadOperationsMetricGroup(),
		NewWorkloadTransactionsMetricGroup(),
		NewWorkloadKeysMetricGroup(),
		NewWorkloadBytesMetricGroup(),
		NewDataMetricGroup(),
	}

	for _, groupToInit := range AllMetricGroups {
		groupToInit.InitScopes()
	}

	return closer
}
