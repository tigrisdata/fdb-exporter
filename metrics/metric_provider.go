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
	"net/http"
	"os"
	"time"

	ulog "github.com/tigrisdata/fdb-exporter/util/log"
)

type MetricProvider struct {
	reporter *MetricReporter
}

func NewMetricProvider() MetricProvider {
	mp := MetricProvider{}
	mp.reporter = NewMetricReporter()
	return mp
}

func (mp *MetricProvider) ServeHttp() {
	listenAddress := os.Getenv("FDB_EXPORTER_HTTP_LISTEN_ADDR")
	if listenAddress == "" {
		listenAddress = DefaultListenAddress
	}
	err := http.ListenAndServe(listenAddress, mp)
	if err != nil {
		ulog.E(err)
		os.Exit(1)
	}
}

func (mp *MetricProvider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mp.reporter.reporter.HTTPHandler().ServeHTTP(w, r)
}

func (m *MetricProvider) Close() {
	ulog.E(m.reporter.closer.Close())
}

// Periodic data collection, called from main in a goroutine
func (mp *MetricProvider) Collect() {
	// TODO make this configurable
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)

	ulog.E(mp.reporter.collectOnce())
	for range ticker.C {
		newReporter := NewMetricReporter()
		ulog.E(newReporter.collectOnce())

		mp.reporter.Close()
		mp.reporter = newReporter
	}
	defer ticker.Stop()
}
