// Copyright © 2025 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

type Config struct {
	// Database is the name of the database to use. A valid database name has the
	// form projects/PROJECT_ID/instances/INSTANCE_ID/databases/DATABASE_ID
	Database string `json:"database" validate:"required"`
	// Endpoint is the URL for endpoint override - testing/dry-run only
	Endpoint string `json:"endpoint"`
}
