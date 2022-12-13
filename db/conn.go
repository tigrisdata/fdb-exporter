package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/tigrisdata/fdb-exporter/models"
)

func getFdb() fdb.Database {
	clusterFile := os.Getenv("FDB_CLUSTER_FILE")
	if clusterFile == "" {
		fmt.Println("Set FDB_CLUSTER_FILE environment variable")
		os.Exit(1)
	}
	fdb.MustAPIVersion(710)
	db, err := fdb.OpenDatabase(clusterFile)
	if err != nil {
		fmt.Printf("Failed to open database using cluster file %s\n", clusterFile)
		os.Exit(1)
	}
	return db
}

func GetStatus() models.FullStatus {
	conn := getFdb()
	var status models.FullStatus
	statusKey := append([]byte{255, 255}, []byte("/status/json")...)
	statusJson, err := conn.ReadTransact(func(t fdb.ReadTransaction) (interface{}, error) {
		return t.Get(fdb.Key(statusKey)).Get()
	})
	if err != nil {
		fmt.Println("Failed to get status")
		os.Exit(1)
	}

	err = json.Unmarshal(statusJson.([]byte), &status)
	if err != nil {
		fmt.Println("Failed to unmarshal status")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return status
}
