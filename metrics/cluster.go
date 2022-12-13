package metrics

import (
	"fmt"

	"github.com/uber-go/tally"
)

var ClusterScope tally.Scope
var WorkloadScope tally.Scope

func InitClusterMetrics() {
	ClusterScope = FdbScope.SubScope("cluster")
	WorkloadScope = ClusterScope.SubScope("workload")
	fmt.Println(ClusterScope)
}
