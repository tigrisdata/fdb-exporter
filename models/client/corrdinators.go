package clientmodel

type Coordinators struct {
	Coordinators    []Coordinator `json:"coordinators"`
	QuorumReachable bool          `json:"quorum_reachable"`
}

type Coordinator struct {
	Address   string `json:"address"`
	Protocol  string `json:"protocol"`
	Reachable bool   `json:"reachable"`
}
