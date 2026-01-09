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
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"strconv"

	"github.com/rs/zerolog/log"
	ulog "github.com/tigrisdata/fdb-exporter/util/log"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/tigrisdata/fdb-exporter/models"
)

const DefaultApiVersion = "710"

var (
	tls     = false
	netopts fdb.NetworkOptions
)

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
	tlsCaFile := os.Getenv("FDB_TLS_CA_FILE")
	tlsCertFile := os.Getenv("FDB_TLS_CERT_FILE")
	tlsKeyFile := os.Getenv("FDB_TLS_KEY_FILE")
	tlsCheckValid := os.Getenv("FDB_TLS_VERIFY_PEERS")

	fdb.MustAPIVersion(apiVersion)
	netopts = fdb.Options()

	// Check TLS in clusterfile
	tls = isTLSmode(&clusterFile)
	if tls {
		// TLS params for fdbclient
		if len(tlsCaFile) > 0 {
			err := netopts.SetTLSCaPath(tlsCaFile)
			if err != nil {
				log.Error().Err(err).Str("msg", "Error cannot set CA").Msg("TLS CA error")
			}
			log.Info().Str("msg", "TLS CA File").Str("fdb.tls-ca-file", tlsCaFile).Msg("TLS CA file set")
		}

		if len(tlsCertFile) > 0 {
			err := netopts.SetTLSCertPath(tlsCertFile)
			if err != nil {
				log.Error().Err(err).Str("msg", "Error cannot set Cert").Msg("TLS Cert error")
			}
			log.Info().Str("msg", "TLS Cert File").Str("fdb.tls-cert-file", tlsCertFile).Msg("TLS cert file set")
		}

		if len(tlsKeyFile) > 0 {
			err := netopts.SetTLSKeyPath(tlsKeyFile)
			if err != nil {
				log.Error().Err(err).Str("msg", "Error cannot set Key").Msg("TLS Key error")
			}
			log.Info().Str("msg", "TLS Private Key").Str("fdb.tls-key-file", tlsKeyFile).Msg("TLS private key set")
		}

		err := netopts.SetTLSVerifyPeers([]byte(tlsCheckValid))
		if err != nil {
			log.Error().Err(err).Str("msg", "Error cannot set VerifyPeers").Msg("TLS VerifyPeers error")
		}
		log.Info().Str("msg", "TLS VerifyPeers").Str("fdb.tls-check-valid", tlsCheckValid).Msg("TLS VerifyPeers set")
	}
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

	if os.Getenv("DEBUG_LOG_INCOMPLETE_STATUS") == "true" {
		if status.Cluster == nil ||
			status.Cluster.DatabaseLockState == nil ||
			status.Cluster.FaultTolerance == nil ||
			status.Cluster.Data == nil {
			log.Debug().Str("status_json", string(statusJson.([]byte))).Msg("status json is missing cluster fields")
		}
	}

	return &status, nil
}

func isTLSmode(c *string) bool {
	// Find if must run in TLS mode
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to open %s", *c)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

	// and then a loop iterates through
	// and prints each of the slice values.
	for _, eachLn := range text {
		//docker:docker@172.19.0.2:4500:tls
		tls, err = regexp.Match(`[0-9]+:tls`, []byte(eachLn))
	}
	return tls
}
