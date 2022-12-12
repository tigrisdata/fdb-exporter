// Copyright 2022 Tigris Data, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
