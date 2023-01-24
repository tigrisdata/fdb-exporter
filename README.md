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

Building
--------

To build a binary, simply run `make`. This will put an `fdb-exporter` binary in the current working directory. To make a container image, run `docker build -t YOUR_ORG_HERE/fdb-exporter:YOUR_TAG_HERE .`.

Testing
-------

To run the tests, use the `make test` command. The tests will start the HTTP server of the interface, and will create the metrics based on bundled json files, and checks for the presence of the metrics. 

Feature requests/bug reports
----------------------------

Feel free to open an issue in github. This exporter is in active development and used in Tigris data. 