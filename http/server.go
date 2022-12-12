package server

import (
	"fmt"
	"github.com/tigrisdata/fdb-exporter/metrics"
	"net/http"
	"os"
)

func Serve() {
	err := http.ListenAndServe(":8080", metrics.Reporter.HTTPHandler())
	if err != nil {
		fmt.Println("Failed to start http listener")
		os.Exit(1)
	}
}
