package clustermodel

type Configuration struct {
	BackupWorkerEnabled            int              `json:"backup_worker_enabled"`
	BlobGranulesEnabled            int              `json:"blob_granules_enabled"`
	CommitProxies                  int              `json:"commit_proxies"`
	CoordinatorsCount              int              `json:"coordinators_count"`
	ExcludedServers                []ExcludedServer `json:"excluded_servers"`
	GrvProxies                     int              `json:"grv_proxies"`
	LogRouters                     int              `json:"log_routers"`
	Logs                           int              `json:"logs"`
	PerpetualStorageWiggle         int              `json:"perpetual_storage_wiggle"`
	PerpetualStorageWiggleLocality string           `json:"perpetual_storage_wiggle_locality""`
	Proxies                        int              `json:"proxies"`
	RedundancyMode                 string           `json:"redundancy_mode"`
	RemoteLogs                     int              `json:"remote_logs"`
	Resolvers                      int              `json:"resolvers"`
	StorageEngine                  string           `json:"storage_engine"`
	StorageMigrationType           string           `json:"storage_migration_type"`
	TenantMode                     string           `json:"tenant_mode"`
	UsableRegions                  int              `json:"usable_regions"`
}

type ExcludedServer struct {
	Address string `json:"address"`
}
