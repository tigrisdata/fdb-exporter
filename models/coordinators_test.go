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

func TestCoordinatorsSingleBasic(t *testing.T) {
	status := CheckJsonFile(t, "status-single-basic.json")
	assert.Equal(t, len(status.Client.Coordinators.Coordinators), 1)
	for _, c := range status.Client.Coordinators.Coordinators {
		assert.Equal(t, c.Address, "127.0.0.1:4689")
		assert.Equal(t, c.Protocol, "0fdb00b071010000")
		assert.True(t, c.Reachable)
	}
}