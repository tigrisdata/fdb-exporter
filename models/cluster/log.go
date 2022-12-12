package clustermodel

type Log struct {
	BeginVersion         int            `json:"begin_version"`
	Current              bool           `json:"current"`
	Epoch                int            `json:"epoch"`
	LogFaultTolerance    int            `json:"log_fault_tolerance"`
	LogInterfaces        []LogInterface `json:"log_interfaces"`
	LogReplicationFactor int            `json:"log_replication_factor"`
	LogWriteAntiQuorum   int            `json:"log_write_anti_quorum"`
	PossiblyLosingData   bool           `json:"possibly_losing_data"`
}

type LogInterface struct {
	Address string `json:"address"`
	Healthy bool   `json:"healthy"`
	Id      string `json:"id"`
}
