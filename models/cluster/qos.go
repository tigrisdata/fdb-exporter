package clustermodel

type Qos struct {
	BatchPerformanceLimitedBy          PerformanceLimitedBy `json:"batch_performance_limited_by"`
	BatchReleasedTransactionsPerSecond float64              `json:"batch_released_transactions_per_second"`
	BatchTransactionsPerSecondLimit    float64              `json:"batch_transactions_per_second_limit"`
	LimitingDataLagStorageServer       Lag                  `json:"limiting_data_lag_storage_server"`
	LimitingDurabilityLagStorageServer Lag                  `json:"limiting_durability_lag_storage_server"`
	LimitingQueueBytesStorageServer    int                  `json:"limiting_queue_bytes_storage_server"`
	PerformanceLimitedBy               PerformanceLimitedBy `json:"performance_limited_by"`
	ReleasedTransactionsPerSecond      float64              `json:"released_transactions_per_second"`
	ThrottledTags                      ThrottledTags        `json:"throttled_tags"`
	TransactionsPerSecondLimit         float64              `json:"transactions_per_second_limit"`
	WorstDataLagStorageServer          Lag                  `json:"worst_data_lag_storage_server"`
	WorstDurabilityLagStorageServer    Lag                  `json:"worst_durability_lag_storage_server"`
	WorstQueueBytesLogServer           int                  `json:"worst_queue_bytes_log_server"`
	WorstQueueBytesStorageServer       int                  `json:"worst_queue_bytes_storage_server"`
}

type PerformanceLimitedBy struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	ReasonId    int    `json:"reason_id"`
}

type ThrottledTags struct {
	Auto   AutoThrottledTags   `json:"auto"`
	Manual ManualThrottledTags `json:"manual"`
}

type AutoThrottledTags struct {
	BusyRead        int `json:"busy_read"`
	BusyWrite       int `json:"busy_write"`
	Count           int `json:"count"`
	RecommendedOnly int `json:"recommended_only"`
}

type ManualThrottledTags struct {
	Count int `json:"count"`
}
