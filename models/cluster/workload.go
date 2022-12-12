package clustermodel

type Workload struct {
	Bytes        WorkloadBytes        `json:"bytes"`
	Keys         WorkloadKeys         `json:"keys"`
	Operations   WorkloadOperations   `json:"operations"`
	Transactions WorkloadTransactions `json:"transactions"`
}

type WorkloadMetrics struct {
	Counter   int     `json:"counter"`
	Hz        float64 `json:"hz"`
	Roughness float64 `json:"roughness"`
}

type WorkloadBytes struct {
	Read    WorkloadMetrics `json:"read"`
	Written WorkloadMetrics `json:"written"`
}

type WorkloadKeys struct {
	Read WorkloadMetrics `json:"read"`
}

type WorkloadOperations struct {
	LocationRequests WorkloadMetrics `json:"location_requests"`
	LowPriorityReads WorkloadMetrics `json:"low_priority_reads"`
	MemoryErrors     WorkloadMetrics `json:"memory_erros"`
	ReadRequests     WorkloadMetrics `json:"read_requests"`
	Reads            WorkloadMetrics `json:"reads"`
	Writes           WorkloadMetrics `json:"writes"`
}

type WorkloadTransactions struct {
	Committed                WorkloadMetrics `json:"committed"`
	Conflicted               WorkloadMetrics `json:"conflicted"`
	RejectedForQueuedTooLong WorkloadMetrics `json:"rejected_for_queued_too_long"`
	Started                  WorkloadMetrics `json:"started"`
	StartedBatchPriority     WorkloadMetrics `json:"started_batch_priority"`
	StartedDefaultPriority   WorkloadMetrics `json:"started_default_priority"`
	StartedImmediatePriority WorkloadMetrics `json:"started_immediate_priority"`
}
