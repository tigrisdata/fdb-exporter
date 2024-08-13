// Copyright 2022-2023 Tigris Data, Inc.
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

package metrics

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	ulog "github.com/tigrisdata/fdb-exporter/util/log"
	"github.com/uber-go/tally"
)

// To include in other structs. Initialization of the map should be done on the included struct.
type scoped struct {
	scopes map[string]tally.Scope
}

func (s *scoped) AddScope(parentScope tally.Scope, key string, name string) {
	s.scopes[key] = parentScope.SubScope(name)
}

func (s *scoped) GetScope(key string) (tally.Scope, error) {
	if scope, ok := s.scopes[key]; ok {
		return scope, nil
	} else {
		return nil, fmt.Errorf("scope %s does not exist", key)
	}
}

func (s *scoped) GetScopeOrExit(key string) tally.Scope {
	scope, err := s.GetScope(key)
	if err != nil {
		msg := "scope not found"
		log.Error().Str("key", key).Msg(msg)
		ulog.E(err, msg)
		os.Exit(1)
	}
	return scope
}

func (s *scoped) CreateScope(key string, name string, parentScope tally.Scope) error {
	err, _ := s.GetScope(key)
	if err == nil {
		return fmt.Errorf("scope %s already exists", key)
	}
	s.scopes[key] = parentScope.SubScope(name)
	return nil
}
