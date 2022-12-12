package clientmodel

type Message struct {
	// Possible names
	// "inconsistent_cluster_file",
	// "unreachable_cluster_controller",
	// "no_cluster_controller",
	// "status_incomplete_client",
	// "status_incomplete_coordinators",
	// "status_incomplete_error",
	// "status_incomplete_timeout",
	// "status_incomplete_cluster",
	// "quorum_not_reachable"
	Name        string
	Description string
}
