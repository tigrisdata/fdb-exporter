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

package db

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
	ulog "github.com/tigrisdata/fdb-exporter/util/log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/tigrisdata/fdb-exporter/models"
)

const DefaultApiVersion = "710"

func getFdb() fdb.Database {
	clusterFile := os.Getenv("FDB_CLUSTER_FILE")
	if clusterFile == "" {
		log.Error().Msg("set FDB_CLUSTER_FILE environment variable")
		os.Exit(1)
	}
	apiVersionStr := os.Getenv("FDB_API_VERSION")
	if apiVersionStr == "" {
		apiVersionStr = DefaultApiVersion
	}
	apiVersion, err := strconv.Atoi(apiVersionStr)
	if err != nil {
		log.Error().Str("FDB_API_VERSION", apiVersionStr).Msg("Could not convert api version to integer")
	}

	fdb.MustAPIVersion(apiVersion)
	db, err := fdb.OpenDatabase(clusterFile)
	if err != nil {
		log.Error().Str("cluster_file", clusterFile).Msg("failed to open database using cluster file")
		os.Exit(1)
	}
	return db
}

func GetStatus() (*models.FullStatus, error) {
	conn := getFdb()
	var status models.FullStatus
	statusKey := append([]byte{255, 255}, []byte("/status/json")...)
	statusJson, err := conn.ReadTransact(func(t fdb.ReadTransaction) (interface{}, error) {
		return t.Get(fdb.Key(statusKey)).Get()
	})
	if err != nil {
		msg := "failed to get status"
		ulog.E(err, msg)
		return nil, err
	}

	err = json.Unmarshal(statusJson.([]byte), &status)
	if err != nil {
		msg := "failed to unmarshal status"
		ulog.E(err, msg)
		return nil, err
	}
	return &status, nil
}
