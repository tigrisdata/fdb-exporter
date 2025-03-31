fdb-exporter
============

Exports a subset of FoundationDB's "status json" command output as prometheus metrics using tally. It uses the FoundationDB client to connect to the cluster. 

Configuration
-------------

This exporter is meant to be run in containers (it can be run a standlone service too, but no init scripts or systemd manifests are provided). The latest container image can be found at `tigrisdata/fdb-exporter:main`. 

The configuration is done through the following environment variables:

* `FDB_CLUSTER_FILE`: the cluster file to connect to FoundationDB
* `FDB_API_VERSION`: the api version to use for the FoundationDB connection (default: 710)
* `ENVIRONMENT`: exposed directly as the `environment` tag in all metrics
* `SERVICE`: exposed directly as the `service` tag in all metrics
* `FDB_VERSION`: exposed directly as the `version` tag in all metrics
* `FDB_CLUSTER_NAME`: exposed directly as the `cluster` tag in all metrics
* `FDB_EXPORTER_HTTP_LISTEN_ADDR`: the address where the built-in web server will listen on and expose the metrics on `/metrics` url. The default is `:8080` 

Optional Variables: 

* `FDB_TLS_CERT_FILE`: the TLS cert file (.crt)  
* `FDB_TLS_KEY_FILE`: the TLS key file (.key)
* `FDB_TLS_VERIFY_PEERS`: the TLS verify peers directive
* `FDB_DEPLOYMENT_NO_K8S`: when set to `""` will make it so the metric tag `fdb_pod_name` will be swapped for `machineid` and the `address` will be set for process specific metrics
* `FDB_EXPORTER_NO_BACKUP_REPORTING`: when set to `""`, the exporter won't report any metrics about backups


Building
--------

To build a binary, simply run `make`. This will put an `fdb-exporter` binary in the current working directory. To make a container image, run `docker build -t YOUR_ORG_HERE/fdb-exporter:YOUR_TAG_HERE .`.

Testing
-------

To run the tests, use the `make test` command. The tests will start the HTTP server of the interface, and will create the metrics based on bundled json files, and checks for the presence of the metrics. 

Feature requests/bug reports
----------------------------

Feel free to open an issue in github. This exporter is in active development and used in Tigris data. 

Metrics
-------
```
# HELP fdb_client_coordinator_quorum fdb_client_coordinator_quorum gauge
# TYPE fdb_client_coordinator_quorum gauge
fdb_client_coordinator_quorum{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_client_coordinator_reachable fdb_client_coordinator_reachable gauge
# TYPE fdb_client_coordinator_reachable gauge
fdb_client_coordinator_reachable{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_client_coordinator_unreachable fdb_client_coordinator_unreachable gauge
# TYPE fdb_client_coordinator_unreachable gauge
fdb_client_coordinator_unreachable{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_client_status_available fdb_client_status_available gauge
# TYPE fdb_client_status_available gauge
fdb_client_status_available{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_client_status_healthy fdb_client_status_healthy gauge
# TYPE fdb_client_status_healthy gauge
fdb_client_status_healthy{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_cluster_backup_config_absent fdb_cluster_backup_config_absent gauge
# TYPE fdb_cluster_backup_config_absent gauge
fdb_cluster_backup_config_absent{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_cluster_commit_latency fdb_cluster_commit_latency gauge
# TYPE fdb_cluster_commit_latency gauge
fdb_cluster_commit_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.5",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.00859928
fdb_cluster_commit_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.95",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.019021
fdb_cluster_commit_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.99",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.019021
# HELP fdb_cluster_data_average_partition_size_bytes fdb_cluster_data_average_partition_size_bytes gauge
# TYPE fdb_cluster_data_average_partition_size_bytes gauge
fdb_cluster_data_average_partition_size_bytes{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 2e+07
# HELP fdb_cluster_data_least_operating_space_bytes_log_server fdb_cluster_data_least_operating_space_bytes_log_server gauge
# TYPE fdb_cluster_data_least_operating_space_bytes_log_server gauge
fdb_cluster_data_least_operating_space_bytes_log_server{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 3.45904433861e+11
# HELP fdb_cluster_data_least_operating_space_bytes_storage_server fdb_cluster_data_least_operating_space_bytes_storage_server gauge
# TYPE fdb_cluster_data_least_operating_space_bytes_storage_server gauge
fdb_cluster_data_least_operating_space_bytes_storage_server{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 9.73707784e+08
# HELP fdb_cluster_data_min_replicas_remaining fdb_cluster_data_min_replicas_remaining gauge
# TYPE fdb_cluster_data_min_replicas_remaining gauge
fdb_cluster_data_min_replicas_remaining{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1
# HELP fdb_cluster_data_missing_data fdb_cluster_data_missing_data gauge
# TYPE fdb_cluster_data_missing_data gauge
fdb_cluster_data_missing_data{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_data_moving_data_in_flight_bytes fdb_cluster_data_moving_data_in_flight_bytes gauge
# TYPE fdb_cluster_data_moving_data_in_flight_bytes gauge
fdb_cluster_data_moving_data_in_flight_bytes{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_data_moving_data_in_queue_bytes fdb_cluster_data_moving_data_in_queue_bytes gauge
# TYPE fdb_cluster_data_moving_data_in_queue_bytes gauge
fdb_cluster_data_moving_data_in_queue_bytes{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_data_moving_data_total_written_types fdb_cluster_data_moving_data_total_written_types gauge
# TYPE fdb_cluster_data_moving_data_total_written_types gauge
fdb_cluster_data_moving_data_total_written_types{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_data_total_disk_used_bytes fdb_cluster_data_total_disk_used_bytes gauge
# TYPE fdb_cluster_data_total_disk_used_bytes gauge
fdb_cluster_data_total_disk_used_bytes{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1.05263104e+08
# HELP fdb_cluster_data_total_kv_size_bytes fdb_cluster_data_total_kv_size_bytes gauge
# TYPE fdb_cluster_data_total_kv_size_bytes gauge
fdb_cluster_data_total_kv_size_bytes{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_grv_latency fdb_cluster_grv_latency gauge
# TYPE fdb_cluster_grv_latency gauge
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="batch",quantile="0.5",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="batch",quantile="0.95",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="batch",quantile="0.99",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="default",quantile="0.5",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.000282764
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="default",quantile="0.95",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.000599384
fdb_cluster_grv_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",priority="default",quantile="0.99",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.000762224
# HELP fdb_cluster_latency_probe_commit_seconds fdb_cluster_latency_probe_commit_seconds gauge
# TYPE fdb_cluster_latency_probe_commit_seconds gauge
fdb_cluster_latency_probe_commit_seconds{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0.0153852
# HELP fdb_cluster_latency_probe_read_seconds fdb_cluster_latency_probe_read_seconds gauge
# TYPE fdb_cluster_latency_probe_read_seconds gauge
fdb_cluster_latency_probe_read_seconds{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0.00018549000000000001
# HELP fdb_cluster_latency_probe_transaction_start_seconds fdb_cluster_latency_probe_transaction_start_seconds gauge
# TYPE fdb_cluster_latency_probe_transaction_start_seconds gauge
fdb_cluster_latency_probe_transaction_start_seconds{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0.003793
# HELP fdb_cluster_latency_probe_transaction_start_seconds_batch fdb_cluster_latency_probe_transaction_start_seconds_batch gauge
# TYPE fdb_cluster_latency_probe_transaction_start_seconds_batch gauge
fdb_cluster_latency_probe_transaction_start_seconds_batch{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0.00029468500000000004
# HELP fdb_cluster_latency_probe_transaction_start_seconds_immediate fdb_cluster_latency_probe_transaction_start_seconds_immediate gauge
# TYPE fdb_cluster_latency_probe_transaction_start_seconds_immediate gauge
fdb_cluster_latency_probe_transaction_start_seconds_immediate{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0.000687361
# HELP fdb_cluster_processes_cpu_cores fdb_cluster_processes_cpu_cores gauge
# TYPE fdb_cluster_processes_cpu_cores gauge
fdb_cluster_processes_cpu_cores{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.0208579
fdb_cluster_processes_cpu_cores{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0.0255029
fdb_cluster_processes_cpu_cores{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.0116077
# HELP fdb_cluster_processes_data_lag_seconds fdb_cluster_processes_data_lag_seconds gauge
# TYPE fdb_cluster_processes_data_lag_seconds gauge
fdb_cluster_processes_data_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1.8114500000000002
fdb_cluster_processes_data_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.8114500000000002
fdb_cluster_processes_data_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.8114500000000002
# HELP fdb_cluster_processes_data_lag_versions fdb_cluster_processes_data_lag_versions gauge
# TYPE fdb_cluster_processes_data_lag_versions gauge
fdb_cluster_processes_data_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1.811449e+06
fdb_cluster_processes_data_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.811449e+06
fdb_cluster_processes_data_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.811449e+06
# HELP fdb_cluster_processes_degraded fdb_cluster_processes_degraded gauge
# TYPE fdb_cluster_processes_degraded gauge
fdb_cluster_processes_degraded{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_degraded{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_degraded{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_disk_busy fdb_cluster_processes_disk_busy gauge
# TYPE fdb_cluster_processes_disk_busy gauge
fdb_cluster_processes_disk_busy{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.0209999
fdb_cluster_processes_disk_busy{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0.0209998
fdb_cluster_processes_disk_busy{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.0209997
# HELP fdb_cluster_processes_disk_free fdb_cluster_processes_disk_free gauge
# TYPE fdb_cluster_processes_disk_free gauge
fdb_cluster_processes_disk_free{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 3.70888474624e+11
fdb_cluster_processes_disk_free{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 3.70888474624e+11
fdb_cluster_processes_disk_free{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 3.70888491008e+11
# HELP fdb_cluster_processes_disk_reads_count fdb_cluster_processes_disk_reads_count gauge
# TYPE fdb_cluster_processes_disk_reads_count gauge
fdb_cluster_processes_disk_reads_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 273037
fdb_cluster_processes_disk_reads_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 273037
fdb_cluster_processes_disk_reads_count{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 273037
# HELP fdb_cluster_processes_disk_reads_hz fdb_cluster_processes_disk_reads_hz gauge
# TYPE fdb_cluster_processes_disk_reads_hz gauge
fdb_cluster_processes_disk_reads_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_disk_reads_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_disk_reads_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_disk_total_bytes fdb_cluster_processes_disk_total_bytes gauge
# TYPE fdb_cluster_processes_disk_total_bytes gauge
fdb_cluster_processes_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 4.99680116736e+11
fdb_cluster_processes_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 4.99680116736e+11
fdb_cluster_processes_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 4.99680116736e+11
# HELP fdb_cluster_processes_disk_writes_count fdb_cluster_processes_disk_writes_count gauge
# TYPE fdb_cluster_processes_disk_writes_count gauge
fdb_cluster_processes_disk_writes_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 5.665219e+06
fdb_cluster_processes_disk_writes_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 5.665219e+06
fdb_cluster_processes_disk_writes_count{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 5.665219e+06
# HELP fdb_cluster_processes_disk_writes_hz fdb_cluster_processes_disk_writes_hz gauge
# TYPE fdb_cluster_processes_disk_writes_hz gauge
fdb_cluster_processes_disk_writes_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 44.7998
fdb_cluster_processes_disk_writes_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 44.7996
fdb_cluster_processes_disk_writes_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 44.7995
# HELP fdb_cluster_processes_durability_bytes fdb_cluster_processes_durability_bytes gauge
# TYPE fdb_cluster_processes_durability_bytes gauge
fdb_cluster_processes_durability_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1934
fdb_cluster_processes_durability_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 82290
fdb_cluster_processes_durability_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1934
# HELP fdb_cluster_processes_durability_lag_seconds fdb_cluster_processes_durability_lag_seconds gauge
# TYPE fdb_cluster_processes_durability_lag_seconds gauge
fdb_cluster_processes_durability_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 5
fdb_cluster_processes_durability_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 5
fdb_cluster_processes_durability_lag_seconds{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 5
# HELP fdb_cluster_processes_durability_lag_versions fdb_cluster_processes_durability_lag_versions gauge
# TYPE fdb_cluster_processes_durability_lag_versions gauge
fdb_cluster_processes_durability_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 5e+06
fdb_cluster_processes_durability_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 5e+06
fdb_cluster_processes_durability_lag_versions{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 5e+06
# HELP fdb_cluster_processes_excluded fdb_cluster_processes_excluded gauge
# TYPE fdb_cluster_processes_excluded gauge
fdb_cluster_processes_excluded{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_excluded{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_excluded{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_input_bytes fdb_cluster_processes_input_bytes gauge
# TYPE fdb_cluster_processes_input_bytes gauge
fdb_cluster_processes_input_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1934
fdb_cluster_processes_input_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 82290
fdb_cluster_processes_input_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1934
# HELP fdb_cluster_processes_kvstore_available_bytes fdb_cluster_processes_kvstore_available_bytes gauge
# TYPE fdb_cluster_processes_kvstore_available_bytes gauge
fdb_cluster_processes_kvstore_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1.073740928e+09
fdb_cluster_processes_kvstore_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.073707784e+09
fdb_cluster_processes_kvstore_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.073740928e+09
# HELP fdb_cluster_processes_kvstore_free_bytes fdb_cluster_processes_kvstore_free_bytes gauge
# TYPE fdb_cluster_processes_kvstore_free_bytes gauge
fdb_cluster_processes_kvstore_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1.073740928e+09
fdb_cluster_processes_kvstore_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.073707784e+09
fdb_cluster_processes_kvstore_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.073740928e+09
# HELP fdb_cluster_processes_kvstore_total_bytes fdb_cluster_processes_kvstore_total_bytes gauge
# TYPE fdb_cluster_processes_kvstore_total_bytes gauge
fdb_cluster_processes_kvstore_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 1.073741824e+09
fdb_cluster_processes_kvstore_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.073741824e+09
fdb_cluster_processes_kvstore_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.073741824e+09
# HELP fdb_cluster_processes_kvstore_used_bytes fdb_cluster_processes_kvstore_used_bytes gauge
# TYPE fdb_cluster_processes_kvstore_used_bytes gauge
fdb_cluster_processes_kvstore_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 94208
fdb_cluster_processes_kvstore_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 192512
fdb_cluster_processes_kvstore_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 94208
# HELP fdb_cluster_processes_log_queue_length fdb_cluster_processes_log_queue_length gauge
# TYPE fdb_cluster_processes_log_queue_length gauge
fdb_cluster_processes_log_queue_length{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
# HELP fdb_cluster_processes_mem_available_bytes fdb_cluster_processes_mem_available_bytes gauge
# TYPE fdb_cluster_processes_mem_available_bytes gauge
fdb_cluster_processes_mem_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 8.589934592e+09
fdb_cluster_processes_mem_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 8.589934592e+09
fdb_cluster_processes_mem_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 8.589934592e+09
# HELP fdb_cluster_processes_mem_limit_bytes fdb_cluster_processes_mem_limit_bytes gauge
# TYPE fdb_cluster_processes_mem_limit_bytes gauge
fdb_cluster_processes_mem_limit_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 8.589934592e+09
fdb_cluster_processes_mem_limit_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 8.589934592e+09
fdb_cluster_processes_mem_limit_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 8.589934592e+09
# HELP fdb_cluster_processes_mem_unused_allocated_memory fdb_cluster_processes_mem_unused_allocated_memory gauge
# TYPE fdb_cluster_processes_mem_unused_allocated_memory gauge
fdb_cluster_processes_mem_unused_allocated_memory{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_mem_unused_allocated_memory{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_mem_unused_allocated_memory{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_mem_unused_bytes fdb_cluster_processes_mem_unused_bytes gauge
# TYPE fdb_cluster_processes_mem_unused_bytes gauge
fdb_cluster_processes_mem_unused_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 2.1997568e+08
fdb_cluster_processes_mem_unused_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1.61775616e+08
fdb_cluster_processes_mem_unused_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 1.77639424e+08
# HELP fdb_cluster_processes_network_conn_closed_hz fdb_cluster_processes_network_conn_closed_hz gauge
# TYPE fdb_cluster_processes_network_conn_closed_hz gauge
fdb_cluster_processes_network_conn_closed_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_network_conn_closed_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_network_conn_closed_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.19999799999999998
# HELP fdb_cluster_processes_network_conn_errors_hz fdb_cluster_processes_network_conn_errors_hz gauge
# TYPE fdb_cluster_processes_network_conn_errors_hz gauge
fdb_cluster_processes_network_conn_errors_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_network_conn_errors_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_network_conn_errors_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_network_conn_established fdb_cluster_processes_network_conn_established gauge
# TYPE fdb_cluster_processes_network_conn_established gauge
fdb_cluster_processes_network_conn_established{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_network_conn_established{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_network_conn_established{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.19999799999999998
# HELP fdb_cluster_processes_network_current_connections fdb_cluster_processes_network_current_connections gauge
# TYPE fdb_cluster_processes_network_current_connections gauge
fdb_cluster_processes_network_current_connections{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 3
fdb_cluster_processes_network_current_connections{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 3
fdb_cluster_processes_network_current_connections{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 4
# HELP fdb_cluster_processes_network_megabits_received fdb_cluster_processes_network_megabits_received gauge
# TYPE fdb_cluster_processes_network_megabits_received gauge
fdb_cluster_processes_network_megabits_received{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.07899489999999999
fdb_cluster_processes_network_megabits_received{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0.129573
fdb_cluster_processes_network_megabits_received{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.0217278
# HELP fdb_cluster_processes_network_megabits_sent fdb_cluster_processes_network_megabits_sent gauge
# TYPE fdb_cluster_processes_network_megabits_sent gauge
fdb_cluster_processes_network_megabits_sent{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0.0993724
fdb_cluster_processes_network_megabits_sent{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0.09080239999999999
fdb_cluster_processes_network_megabits_sent{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0.0391868
# HELP fdb_cluster_processes_query_count fdb_cluster_processes_query_count gauge
# TYPE fdb_cluster_processes_query_count gauge
fdb_cluster_processes_query_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_query_count{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 1716
fdb_cluster_processes_query_count{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_query_hz fdb_cluster_processes_query_hz gauge
# TYPE fdb_cluster_processes_query_hz gauge
fdb_cluster_processes_query_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_query_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 6.39984
fdb_cluster_processes_query_hz{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_queue_disk_available_bytes fdb_cluster_processes_queue_disk_available_bytes gauge
# TYPE fdb_cluster_processes_queue_disk_available_bytes gauge
fdb_cluster_processes_queue_disk_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_queue_disk_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_queue_disk_available_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_queue_disk_free_bytes fdb_cluster_processes_queue_disk_free_bytes gauge
# TYPE fdb_cluster_processes_queue_disk_free_bytes gauge
fdb_cluster_processes_queue_disk_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_queue_disk_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_queue_disk_free_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_queue_disk_total_bytes fdb_cluster_processes_queue_disk_total_bytes gauge
# TYPE fdb_cluster_processes_queue_disk_total_bytes gauge
fdb_cluster_processes_queue_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_queue_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_queue_disk_total_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_queue_disk_used_bytes fdb_cluster_processes_queue_disk_used_bytes gauge
# TYPE fdb_cluster_processes_queue_disk_used_bytes gauge
fdb_cluster_processes_queue_disk_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_queue_disk_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_queue_disk_used_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_processes_stored_bytes fdb_cluster_processes_stored_bytes gauge
# TYPE fdb_cluster_processes_stored_bytes gauge
fdb_cluster_processes_stored_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_processes_stored_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0
fdb_cluster_processes_stored_bytes{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_read_latency fdb_cluster_read_latency gauge
# TYPE fdb_cluster_read_latency gauge
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",quantile="0.5",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",quantile="0.95",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="1",log="1",master="0",quantile="0.99",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="be7dacbc31d0"} 0
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",quantile="0.5",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 5.17368e-05
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",quantile="0.95",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 9.274480000000001e-05
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="0",commit_proxy="0",coordinator="0",data_distributor="1",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="1",quantile="0.99",ratekeeper="1",resolver="1",service="default_service",storage="1",version="default_version",zone="b6ef736d9565"} 0.00011897100000000001
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.5",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.95",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
fdb_cluster_read_latency{class_type="unset",cluster="default_cluster_name",cluster_controller="1",commit_proxy="1",coordinator="1",data_distributor="0",env="default_env",fdb_cluster="default_fdb_cluster_name",fdb_pod_name="",grv_proxy="0",log="0",master="0",quantile="0.99",ratekeeper="0",resolver="0",service="default_service",storage="1",version="default_version",zone="99dcaaf374f0"} 0
# HELP fdb_cluster_workload_bytes_read_count fdb_cluster_workload_bytes_read_count gauge
# TYPE fdb_cluster_workload_bytes_read_count gauge
fdb_cluster_workload_bytes_read_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 903223
# HELP fdb_cluster_workload_bytes_read_hz fdb_cluster_workload_bytes_read_hz gauge
# TYPE fdb_cluster_workload_bytes_read_hz gauge
fdb_cluster_workload_bytes_read_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 3903.1
# HELP fdb_cluster_workload_bytes_written_count fdb_cluster_workload_bytes_written_count gauge
# TYPE fdb_cluster_workload_bytes_written_count gauge
fdb_cluster_workload_bytes_written_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 6428
# HELP fdb_cluster_workload_bytes_written_hz fdb_cluster_workload_bytes_written_hz gauge
# TYPE fdb_cluster_workload_bytes_written_hz gauge
fdb_cluster_workload_bytes_written_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_keys_read_count fdb_cluster_workload_keys_read_count gauge
# TYPE fdb_cluster_workload_keys_read_count gauge
fdb_cluster_workload_keys_read_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 2913
# HELP fdb_cluster_workload_keys_read_hz fdb_cluster_workload_keys_read_hz gauge
# TYPE fdb_cluster_workload_keys_read_hz gauge
fdb_cluster_workload_keys_read_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 12.1997
# HELP fdb_cluster_workload_operations_location_requests_count fdb_cluster_workload_operations_location_requests_count gauge
# TYPE fdb_cluster_workload_operations_location_requests_count gauge
fdb_cluster_workload_operations_location_requests_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 42
# HELP fdb_cluster_workload_operations_location_requests_hz fdb_cluster_workload_operations_location_requests_hz gauge
# TYPE fdb_cluster_workload_operations_location_requests_hz gauge
fdb_cluster_workload_operations_location_requests_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_operations_low_priority_reads_count fdb_cluster_workload_operations_low_priority_reads_count gauge
# TYPE fdb_cluster_workload_operations_low_priority_reads_count gauge
fdb_cluster_workload_operations_low_priority_reads_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_operations_low_priority_reads_hz fdb_cluster_workload_operations_low_priority_reads_hz gauge
# TYPE fdb_cluster_workload_operations_low_priority_reads_hz gauge
fdb_cluster_workload_operations_low_priority_reads_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_operations_memory_errors_count fdb_cluster_workload_operations_memory_errors_count gauge
# TYPE fdb_cluster_workload_operations_memory_errors_count gauge
fdb_cluster_workload_operations_memory_errors_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_operations_memory_errors_hz fdb_cluster_workload_operations_memory_errors_hz gauge
# TYPE fdb_cluster_workload_operations_memory_errors_hz gauge
fdb_cluster_workload_operations_memory_errors_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_operations_read_requests_count fdb_cluster_workload_operations_read_requests_count gauge
# TYPE fdb_cluster_workload_operations_read_requests_count gauge
fdb_cluster_workload_operations_read_requests_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1716
# HELP fdb_cluster_workload_operations_read_requests_hz fdb_cluster_workload_operations_read_requests_hz gauge
# TYPE fdb_cluster_workload_operations_read_requests_hz gauge
fdb_cluster_workload_operations_read_requests_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 6.39984
# HELP fdb_cluster_workload_operations_reads_count fdb_cluster_workload_operations_reads_count gauge
# TYPE fdb_cluster_workload_operations_reads_count gauge
fdb_cluster_workload_operations_reads_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 1716
# HELP fdb_cluster_workload_operations_reads_hz fdb_cluster_workload_operations_reads_hz gauge
# TYPE fdb_cluster_workload_operations_reads_hz gauge
fdb_cluster_workload_operations_reads_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 6.39984
# HELP fdb_cluster_workload_operations_writes_count fdb_cluster_workload_operations_writes_count gauge
# TYPE fdb_cluster_workload_operations_writes_count gauge
fdb_cluster_workload_operations_writes_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 71
# HELP fdb_cluster_workload_operations_writes_hz fdb_cluster_workload_operations_writes_hz gauge
# TYPE fdb_cluster_workload_operations_writes_hz gauge
fdb_cluster_workload_operations_writes_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_transactions_committed_count fdb_cluster_workload_transactions_committed_count gauge
# TYPE fdb_cluster_workload_transactions_committed_count gauge
fdb_cluster_workload_transactions_committed_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 41
# HELP fdb_cluster_workload_transactions_committed_hz fdb_cluster_workload_transactions_committed_hz gauge
# TYPE fdb_cluster_workload_transactions_committed_hz gauge
fdb_cluster_workload_transactions_committed_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_transactions_conflicted_count fdb_cluster_workload_transactions_conflicted_count gauge
# TYPE fdb_cluster_workload_transactions_conflicted_count gauge
fdb_cluster_workload_transactions_conflicted_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 5
# HELP fdb_cluster_workload_transactions_conflicted_hz fdb_cluster_workload_transactions_conflicted_hz gauge
# TYPE fdb_cluster_workload_transactions_conflicted_hz gauge
fdb_cluster_workload_transactions_conflicted_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_transactions_default_count fdb_cluster_workload_transactions_default_count gauge
# TYPE fdb_cluster_workload_transactions_default_count gauge
fdb_cluster_workload_transactions_default_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="default",service="default_service",version="default_version"} 428
# HELP fdb_cluster_workload_transactions_default_hz fdb_cluster_workload_transactions_default_hz gauge
# TYPE fdb_cluster_workload_transactions_default_hz gauge
fdb_cluster_workload_transactions_default_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="default",service="default_service",version="default_version"} 1.39999
# HELP fdb_cluster_workload_transactions_immediate_count fdb_cluster_workload_transactions_immediate_count gauge
# TYPE fdb_cluster_workload_transactions_immediate_count gauge
fdb_cluster_workload_transactions_immediate_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="immediate",service="default_service",version="default_version"} 278
# HELP fdb_cluster_workload_transactions_immediate_hz fdb_cluster_workload_transactions_immediate_hz gauge
# TYPE fdb_cluster_workload_transactions_immediate_hz gauge
fdb_cluster_workload_transactions_immediate_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="immediate",service="default_service",version="default_version"} 0.999991
# HELP fdb_cluster_workload_transactions_rejected_for_queued_too_long_count fdb_cluster_workload_transactions_rejected_for_queued_too_long_count gauge
# TYPE fdb_cluster_workload_transactions_rejected_for_queued_too_long_count gauge
fdb_cluster_workload_transactions_rejected_for_queued_too_long_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz gauge
# TYPE fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz gauge
fdb_cluster_workload_transactions_rejected_for_queued_too_long_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",service="default_service",version="default_version"} 0
# HELP fdb_cluster_workload_transactions_started_count fdb_cluster_workload_transactions_started_count gauge
# TYPE fdb_cluster_workload_transactions_started_count gauge
fdb_cluster_workload_transactions_started_count{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="batch",service="default_service",version="default_version"} 1
# HELP fdb_cluster_workload_transactions_started_hz fdb_cluster_workload_transactions_started_hz gauge
# TYPE fdb_cluster_workload_transactions_started_hz gauge
fdb_cluster_workload_transactions_started_hz{cluster="default_cluster_name",env="default_env",fdb_cluster="default_fdb_cluster_name",priority="batch",service="default_service",version="default_version"} 0
```