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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/tigrisdata/fdb-exporter/db"

	"github.com/tigrisdata/fdb-exporter/models"
	ulog "github.com/tigrisdata/fdb-exporter/util/log"
	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
)

const RelativeJsonFileLocation = "test/data"

type MetricInfo struct {
	groups   []Collectable
	status   *models.FullStatus
	closer   io.Closer
	reporter promreporter.Reporter
	scoped
}

type fetchedMetric struct {
	metricKey   string
	metricValue string
	tags        string
}

func NewMetricInfo() MetricInfo {
	m := MetricInfo{}
	m.scopes = make(map[string]tally.Scope)

	m.reporter = promreporter.NewReporter(promreporter.Options{})
	m.scopes["root"], m.closer = tally.NewRootScope(tally.ScopeOptions{
		Tags:           GetBaseTags(),
		CachedReporter: m.reporter,
		Separator:      promreporter.DefaultSeparator,
	}, 1*time.Second)

	// TODO check if the scopes actually exist
	m.scopes["fdb"] = m.scopes["root"].SubScope("fdb")
	m.scopes["client"] = m.scopes["fdb"].SubScope("client")
	m.scopes["cluster"] = m.scopes["fdb"].SubScope("cluster")
	m.scopes["workload"] = m.scopes["cluster"].SubScope("workload")

	m.groups = []Collectable{
		NewCoordinatorMetricGroup(&m),
		NewDbStatusMetricGroup(&m),
		NewWorkloadOperationsMetricGroup(&m),
		NewWorkloadTransactionsMetricGroup(&m),
		NewWorkloadKeysMetricGroup(&m),
		NewWorkloadBytesMetricGroup(&m),
		NewDataMetricGroup(&m),
	}

	return m
}

func (m *MetricInfo) Collect() {
	// TODO make this configurable
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	ulog.E(m.collectOnce())
	for range ticker.C {
		ulog.E(m.collectOnce())
	}
	defer ticker.Stop()
}

func (m *MetricInfo) collectOnce() error {
	var err error
	m.status, err = db.GetStatus()
	if err != nil {
		return fmt.Errorf("failed to get status")
	}

	if len(m.groups) == 0 {
		ulog.E(fmt.Errorf("no metric groups detected"))
	}

	for _, group := range m.groups {
		group.GetMetrics(m.status)
	}
	log.Debug().Msg("collected once")
	return nil
}

func (m *MetricInfo) collectOnceFromFile(fileName string) error {
	// Used in testing
	wd, err := os.Getwd()
	if err != nil {
		ulog.E(err)
	}
	testFilePath := fmt.Sprintf("%s/../%s/%s", wd, RelativeJsonFileLocation, fileName)
	f, err := os.Open(testFilePath)
	if err != nil {
		ulog.E(err)
	}
	defer f.Close()
	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		ulog.E(err)
	}
	err = json.Unmarshal(jsonBytes, &m.status)
	if err != nil {
		ulog.E(err)
	}

	for _, group := range m.groups {
		group.GetMetrics(m.status)
	}
	return nil
}

func (m *MetricInfo) Close() {
	ulog.E(m.closer.Close())
}

func (m *MetricInfo) ServeHttp() {
	err := http.ListenAndServe(":8080", m.reporter.HTTPHandler())
	if err != nil {
		ulog.E(err)
		os.Exit(1)
	}
}
