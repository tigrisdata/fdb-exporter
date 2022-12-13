package main

import (
	"sync"

	server "github.com/tigrisdata/fdb-exporter/http"
	"github.com/tigrisdata/fdb-exporter/metrics"
)

func main() {
	var wg sync.WaitGroup
	closer := metrics.InitMetrics()

	defer closer.Close()

	go server.Serve()
	wg.Add(1)

	go metrics.Collect()
	wg.Add(1)

	wg.Wait()
}
