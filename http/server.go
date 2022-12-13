package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tigrisdata/fdb-exporter/metrics"
)

func Serve() {
	err := http.ListenAndServe(":8080", metrics.Reporter.HTTPHandler())
	if err != nil {
		fmt.Println("Failed to start http listener")
		os.Exit(1)
	}
}
