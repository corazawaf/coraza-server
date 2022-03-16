// Copyright 2022 The Corazawaf Authors.
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

package engines

import (
	"errors"
)

// Engine is the interface that all protocols must implement.
type Engine interface {
	Init() error
	Start() error
}

// Get returns the instance of the engine with the given name.
func Get(name string) (Engine, error) {
	switch name {
	case "http":
		return nil, nil
	case "grpc":
		return nil, nil
	case "uds":
		return nil, nil
	default:
		return nil, errors.New("the protocol is not supported in the coraza-server")
	}
}
