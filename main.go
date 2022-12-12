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

package main

import (
	"fmt"
	server "github.com/tigrisdata/fdb-exporter/http"
	"github.com/tigrisdata/fdb-exporter/metrics"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	closer := metrics.InitMetrics()

	go server.Serve()
	wg.Add(1)

	go metrics.Collect()
	wg.Add(1)

	wg.Wait()
	if closer != nil {
		err := closer.Close()
		if err != nil {
			fmt.Println("failed to close")
			os.Exit(1)
		}
	}
}
