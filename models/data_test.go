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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, status.Cluster.Data.AveragePartitionSizeBytes, 20668340)
	assert.Equal(t, status.Cluster.Data.LeastOperatingSpaceBytesLogServer, 604285174979)
	assert.Equal(t, status.Cluster.Data.LeastOperatingSpaceBytesStorageServer, 854875383)
	assert.Equal(t, status.Cluster.Data.PartitionsCount, 2)

}
