package clientmodel

// Client top level status
type Status struct {
	ClusterFile    ClusterFile    `json:"cluster_file"`
	Coordinators   Coordinators   `json:"coordinators"`
	DatabaseStatus DatabaseStatus `json:"database_status"`
	Messages       []Message      `json:"messages"`
	Timestamp      int            `json:"timestamp"`
}
