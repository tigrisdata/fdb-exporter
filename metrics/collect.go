package metrics

import (
	"context"
	"fmt"
	"github.com/tigrisdata/fdb-exporter/db"
	"os"
	"time"
)

func Collect() {
	// TODO make this configurable
	interval := 10 * time.Second
	ticker := time.NewTicker(interval)
	// First collection should happen immediately, the rest is timer based
	ctx, cancel := context.WithTimeout(context.Background(), interval)
	defer cancel()
	collectOnce(ctx)
	for range ticker.C {
		collectOnce(ctx)
	}
	defer ticker.Stop()
}

func collectOnce(ctx context.Context) {
	status := db.GetStatus()

	if len(AllMetricGroups) == 0 {
		fmt.Println("No metric groups detected, exiting")
		os.Exit(1)
	}

	for _, group := range AllMetricGroups {
		group.SetStatus(&status)
		group.GetMetrics()
	}

	fmt.Println("collected once")
}
