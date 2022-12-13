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
	"io"
	"time"

	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
)

var RootScope tally.Scope
var FdbScope tally.Scope
var Reporter promreporter.Reporter
var AllMetricGroups []Collectable

func InitMetrics() io.Closer {
	// var closer io.Closer
	Reporter = promreporter.NewReporter(promreporter.Options{})
	RootScope, closer := tally.NewRootScope(tally.ScopeOptions{
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
