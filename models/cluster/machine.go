package clustermodel

type Machine struct {
	Address             string          `json:"address"`
	ContributingWorkers int             `json:"contributing_workers"`
	Cpu                 MachineCpu      `json:"cpu"`
	Excluded            bool            `json:"excluded"`
	Locality            MachineLocality `json:"locality"`
	Memory              MachineMemory   `json:"memory"`
	Network             MachineNetwork  `json:"network"`
}

type MachineCpu struct {
	LogicalCoreUtilization float64 `json:"logical_core_utilization"`
}

type MachineLocality struct {
	MachineId string `json:"machine_id"`
	ProcessId string `json:"process_id"`
	ZoneId    string `json:"zone_id"`
}

type MachineMemory struct {
	CommittedBytes int `json:"committed_bytes"`
	FreeBytes      int `json:"free_bytes"`
	TotalBytes     int `json:"total_bytes"`
}

type MachineNetwork struct {
	MegabitsReceived         Hz `json:"megabits_received"`
	MegabitsSent             Hz `json:"megabits_sent"`
	TcpSegmentsRetransmitted Hz `json:"tcp_segments_retransmitted"`
}
