package metrics

import (
	"github.com/tigrisdata/fdb-exporter/models"
	"github.com/uber-go/tally"
)

// TODO: finish this
type DataMetricGroup struct {
	MetricGroup
}

func NewDataMetricGroup() *DataMetricGroup {
	d := DataMetricGroup{}
	d.scopes = make(map[string]tally.Scope)
	return &d
}

func (d *DataMetricGroup) SetStatus(status *models.FullStatus) {
	d.status = status
}

func (d *DataMetricGroup) InitScopes() {
	d.scopes["default"] = ClusterScope.SubScope("data")
}

func (d *DataMetricGroup) GetMetrics() {
	metrics := map[string]int{
		"average_partition_size_bytes":               d.status.Cluster.Data.AveragePartitionSizeBytes,
		"least_operating_space_bytes_log_server":     d.status.Cluster.Data.LeastOperatingSpaceBytesLogServer,
		"least_operating_space_bytes_storage_server": d.status.Cluster.Data.LeastOperatingSpaceBytesStorageServer,
		"moving_data_in_flight_bytes":                d.status.Cluster.Data.MovingData.InFlightBytes,
		"moving_data_in_queue_bytes":                 d.status.Cluster.Data.MovingData.InQueueBytes,
		"moving_data_total_written_types":            d.status.Cluster.Data.MovingData.TotalWrittenBytes,
		"total_disk_used_bytes":                      d.status.Cluster.Data.TotalDiskUsedBytes,
		"total_kv_size_bytes":                        d.status.Cluster.Data.TotalKvSizeBytes,
	}
	for name, value := range metrics {
		SetIntGauge(d.scopes["default"], name, GetBaseTags(), value)
	}
}
