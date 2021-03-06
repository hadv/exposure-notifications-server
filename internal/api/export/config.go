// Copyright 2020 Google LLC
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

package export

import (
	"time"

	"github.com/google/exposure-notifications-server/internal/database"
	"github.com/google/exposure-notifications-server/internal/setup"
)

// Compile-time check to assert this config matches requirements.
var _ setup.KeyManagerProvider = (*Config)(nil)
var _ setup.BlobStorageConfigProvider = (*Config)(nil)
var _ setup.DBConfigProvider = (*Config)(nil)

// Config represents the configuration and associated environment variables for
// the export components.
type Config struct {
	Database          *database.Config
	Port              string        `envconfig:"PORT" default:"8080"`
	CreateTimeout     time.Duration `envconfig:"CREATE_BATCHES_TIMEOUT" default:"5m"`
	WorkerTimeout     time.Duration `envconfig:"WORKER_TIMEOUT" default:"5m"`
	MinRecords        int           `envconfig:"EXPORT_FILE_MIN_RECORDS" default:"1000"`
	PaddingRange      int           `envconfig:"EXPORT_FILE_PADDING_RANGE" default:"100"`
	MaxRecords        int           `envconfig:"EXPORT_FILE_MAX_RECORDS" default:"30000"`
	DefaultKeyID      string        `envconfig:"EXPORT_FILE_DEFAULT_KEY_ID" default:"ExampleServer"`
	DefaultKeyVersion string        `envconfig:"EXPORT_FILE_DEFAULT_KEY_VERSION" default:"1"`
}

// DB returns the database config.
func (c *Config) DB() *database.Config {
	return c.Database
}

// KeyManager returns the KeyManager configuration.
func (c *Config) KeyManager() bool {
	return true
}

// BlobStorage returns the BlobStorage configuration.
func (c *Config) BlobStorage() bool {
	return true
}
