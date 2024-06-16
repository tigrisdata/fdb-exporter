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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/tigrisdata/fdb-exporter/db"

	"github.com/tigrisdata/fdb-exporter/models"
	ulog "github.com/tigrisdata/fdb-exporter/util/log"
	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
)

const (
	RelativeJsonFileLocation = "test/data"
	DefaultListenAddress     = ":8080"
)

// High level type that can report all the metrics.
// Reponsible for:
// * getting the status json output periodically
// * call the GetMetrics() method of each group, which will set up the metrics in tally
type MetricReporter struct {
	groups   []Collectable
	status   *models.FullStatus
	closer   io.Closer
	reporter promreporter.Reporter
	scoped
}

// Used in integration testing when the metrics are fetched from the server
type fetchedMetric struct {
	key   string
	value string
	tags  string
}

func NewMetricReporter() *MetricReporter {
	m := MetricReporter{}
	m.scopes = make(map[string]tally.Scope)

	m.reporter = promreporter.NewReporter(promreporter.Options{})
	m.scopes["root"], m.closer = tally.NewRootScope(tally.ScopeOptions{
		Tags:           GetBaseTags(),
		CachedReporter: m.reporter,
		Separator:      promreporter.DefaultSeparator,
	}, 1*time.Second)

	m.AddScope(m.scopes["root"], "fdb", "fdb")
	m.AddScope(m.scopes["fdb"], "client", "client")
	m.AddScope(m.scopes["fdb"], "cluster", "cluster")
	m.AddScope(m.scopes["cluster"], "workload", "workload")

	// Add each impltemented group here
	m.groups = []Collectable{
		NewCoordinatorMetricGroup(&m),
		NewDbStatusMetricGroup(&m),
		NewWorkloadOperationsMetricGroup(&m),
		NewWorkloadTransactionsMetricGroup(&m),
		NewWorkloadKeysMetricGroup(&m),
		NewWorkloadBytesMetricGroup(&m),
		NewDataMetricGroup(&m),
		NewProcessesMetricGroup(&m),
		NewLatencyProbeMetricGroup(&m),
		NewBackupMetricGroup(&m),
		NewClusterMessageMetricGroup(&m),
	}
	return &m
}

// Periodic data collection, called from main in a goroutine
func (m *MetricReporter) Collect() {
	// TODO make this configurable
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)

	ulog.E(m.collectOnce())
	for range ticker.C {
		ulog.E(m.collectOnce())
	}
	defer ticker.Stop()
}

// Single data collection, fetches status, gets the metrics from each group
func (m *MetricReporter) collectOnce() error {
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
	return nil
}

// Used only in integration testing, collects metrics from a json file
func (m *MetricReporter) collectOnceFromFile(fileName string) error {
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
	jsonBytes, err := io.ReadAll(f)
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

func (m *MetricReporter) Close() {
	ulog.E(m.closer.Close())
}

func (m *MetricReporter) ServeHttp() {
	listenAddress := os.Getenv("FDB_EXPORTER_HTTP_LISTEN_ADDR")
	if listenAddress == "" {
		listenAddress = DefaultListenAddress
	}
	err := http.ListenAndServe(listenAddress, m.reporter.HTTPHandler())
	if err != nil {
		ulog.E(err)
		os.Exit(1)
	}
}
