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
