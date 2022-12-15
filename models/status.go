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

package models

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

// Top level fields from status json
type FullStatus struct {
	Client  ClientStatus  `json:"client"`
	Cluster ClusterStatus `json:"cluster"`
}

const RelativeJsonFileLocation = "test/data"

func GetStatusFromFile(fileName string) (*FullStatus, error) {
	var status FullStatus
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory")
	}
	testFilePath := fmt.Sprintf("%s/../%s/%s", wd, RelativeJsonFileLocation, fileName)
	f, err := os.Open(testFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open test file %s", testFilePath)
	}
	defer f.Close()

	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read test file %s", testFilePath)
	}
	err = json.Unmarshal(jsonBytes, &status)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshel test file %s", testFilePath)
	}
	return &status, nil
}

func CheckJsonFile(t *testing.T, fileName string) *FullStatus {
	status, err := GetStatusFromFile(fileName)
	assert.Nil(t, err, "error reading status file")
	return status
}
