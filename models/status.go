package models

import (
	clientmodel "github.com/tigrisdata/fdb-exporter/models/client"
	clustermodel "github.com/tigrisdata/fdb-exporter/models/cluster"
)

// Top level fields from status json
type FullStatus struct {
	Client  clientmodel.Status  `json:"client"`
	Cluster clustermodel.Status `json:"cluster"`
}
