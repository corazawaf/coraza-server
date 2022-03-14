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

package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

// C is used to store the configuration.
var C Config

func init() {
	flag.StringVar(&C.ConfigFile, "config-file", "./config.yml", "The configuration file of the coraza-server. (default: ./config.yml)")
}

// Config is used to configure coraza-server.
type Config struct {
	Log Log `yaml:"log"`

	// ConfigFile is the configuration file of the coraza-server.
	ConfigFile string
}

// Log is used to configure the level and dir of the log.
type Log struct {
	Level string `yaml:"level"`
	Dir   string `yaml:"dir"`
}

// InitConfig initializes the configuration.
func InitConfig() error {
	f, err := os.Open(C.ConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&C)
	if err != nil {
		return err
	}

	return nil
}
