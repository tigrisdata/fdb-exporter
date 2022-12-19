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
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getMetricsFromTestFile(t *testing.T, fileName string) []fetchedMetric {
	var res []fetchedMetric
	m := NewMetricInfo()
	ts := httptest.NewServer(m.reporter.HTTPHandler())
	defer ts.Close()
	err := m.collectOnceFromFile(fileName)
	if err != nil {
		assert.Nil(t, err, "error collecting metrics from file %s", fileName)
	}
	resGet, err := http.Get(ts.URL)
	if err != nil {
		assert.Nil(t, err, "error getting metrics")
	}
	out, err := ioutil.ReadAll(resGet.Body)
	if err != nil {
		assert.Nil(t, err, "error reading metrics")
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		res = append(res, parseMetricLine(t, line))
	}
	return res
}

func parseMetricLine(t *testing.T, line string) fetchedMetric {
	fullMetricRe, err := regexp.Compile(`(?P<name>[a-z_]+)\{(?P<tags>.*)\} (?P<value>[a-zA-Z0-9]+)`)

	assert.Nil(t, err, "error parsing metric line with regexp")
	match := fullMetricRe.FindStringSubmatch(line)
	if len(match) == 4 {
		return fetchedMetric{metricKey: match[1], tags: match[2], metricValue: match[3]}
	} else {
		return fetchedMetric{}
	}
}

func checkMetrics(t *testing.T, metrics []fetchedMetric, expectedMetrics []string) {
	for _, expectedMetric := range expectedMetrics {
		foundExpectedMetric := false
		for _, metric := range metrics {
			if metric.metricKey == expectedMetric && metric.metricValue != "" {
				foundExpectedMetric = true
			}
		}
		if !foundExpectedMetric {
			assert.Falsef(t, true, "did not find %s in the metrics", expectedMetric)
		}
		assert.True(t, foundExpectedMetric)
	}
}
