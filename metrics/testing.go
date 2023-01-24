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
	"bufio"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMetricReporter *MetricReporter

func getMetricsFromTestFile(t *testing.T, fileName string) []fetchedMetric {
	var res []fetchedMetric
	ts := httptest.NewServer(testMetricReporter.reporter.HTTPHandler())
	defer ts.Close()
	err := testMetricReporter.collectOnceFromFile(fileName)
	if err != nil {
		assert.Nil(t, err, "error collecting metrics from file %s", fileName)
	}
	testMetricReporter.reporter.Flush()
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
		return fetchedMetric{key: match[1], tags: match[2], value: match[3]}
	} else {
		return fetchedMetric{}
	}
}

func checkMetrics(t *testing.T, metrics []fetchedMetric, expectedMetrics []string) {
	for _, expectedMetric := range expectedMetrics {
		foundExpectedMetric := false
		for _, metric := range metrics {
			if metric.key == expectedMetric && metric.value != "" {
				foundExpectedMetric = true
			}
		}
		if !foundExpectedMetric {
			assert.Falsef(t, true, "did not find %s in the metrics", expectedMetric)
		}
		assert.True(t, foundExpectedMetric)
	}
}

func checkTagsForMetric(t *testing.T, metrics []fetchedMetric, metricToCheck string, tagsToCheck []string) {
	metricFound := false
	for _, metric := range metrics {
		if metric.key == metricToCheck {
			metricFound = true
			for _, tagToCheck := range tagsToCheck {
				if !strings.Contains(metric.tags, tagToCheck) {
					assert.Fail(t, "metric key %s tags %s value %s, did not contain tag %s", metric.key, metric.tags, metric.value, tagToCheck)
				}
			}
		}
	}
	if !metricFound {
		assert.Fail(t, "metric %s not found", metricToCheck)
	}
}

func initTestMetricReporter() {
	if testMetricReporter == nil {
		testMetricReporter = NewMetricReporter()
	}
}
