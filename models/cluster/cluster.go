package clustermodel

// Cluster top level status
type Status struct {
	// TODO: add support for clients and detecting version compatibility
	// TODO: add support for incompatible connections
	// TODO: add layers
	ClusterControllerTimestamp int                `json:"cluster_controller_timestamp"`
	Configuration              Configuration      `json:"configuration"`
	ConnectionString           string             `json:"connection_string"`
	Data                       Data               `json:"data"`
	DatabaseAvailable          bool               `json:"database_available"`
	DatabaseLockState          LockState          `json:"database_lock_state"`
	DatacenterLag              Lag                `json:"datacenter_lag"`
	DegradedProcesses          int                `json:"degraded_processes"`
	FaultTolerance             FaultTolerance     `json:"fault_tolerance"`
	FullReplication            bool               `json:"full_replication"`
	Generation                 int                `json:"generation"`
	LatencyProbe               LatencyProbe       `json:"latency_probe"`
	Logs                       []Log              `json:"logs"`
	Machines                   map[string]Machine `json:"machines"`
	Messages                   []Message          `json:"messages"`
	PageCache                  PageCache          `json:"page_cache"`
	Processes                  map[string]Process `json:"processes"`
	ProtocolVersion            string             `json:"protocol_version"`
	Qos                        Qos                `json:"qos"`
	RecoveryState              RecoveryState      `json:"recovery_state"`
	Workload                   Workload           `json:"workload"`
}

type LockState struct {
	Locked bool `json:"locked"`
}

type Lag struct {
	Seconds  float64 `json:"seconds"`
	Versions int     `json:"versions"`
}

type FaultTolerance struct {
	MaxZoneFailuresWithoutLosingAvailability int `json:"max_zone_failures_without_losing_availability"`
	MaxZoneFailuresWithoutLosingData         int `json:"max_zone_failures_without_losing_data"`
}

type LatencyProbe struct {
	BatchPriorityTransactionStartSeconds     float64 `json:"batch_priority_transaction_start_seconds"`
	CommitSeconds                            float64 `json:"commit_seconds"`
	ImmediatePriorityTransactionStartSeconds float64 `json:"immediate_priority_transaction_start_seconds"`
	ReadSeconds                              float64 `json:"read_seconds"`
	TransactionStartSeconds                  float64 `json:"transaction_start_seconds"`
}

type Hz struct {
	Hz float64 `json:"hz"`
}

type PageCache struct {
	LogHitRate     int `json:"log_hit_rate"`
	StorageHitRate int `json:"storage_hit_rate"`
}

type RecoveryState struct {
	ActiveGenerations         int     `json:"active_generations"`
	Description               string  `json:"description"`
	Name                      string  `json:"name"`
	SecondsSinceLastRecovered float64 `json:"seconds_since_last_recovered"`
}
