// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package destination

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	ConfigDatabase = "database"
	ConfigEndpoint = "endpoint"
	ConfigTable    = "table"
)

func (Config) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		ConfigDatabase: {
			Default:     "",
			Description: "Database is the name of the database to use. A valid database name has the\nform projects/PROJECT_ID/instances/INSTANCE_ID/databases/DATABASE_ID",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigEndpoint: {
			Default:     "",
			Description: "Endpoint is the URL for endpoint override - testing/dry-run only",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigTable: {
			Default:     "{{ index .Metadata \"opencdc.collection\" }}",
			Description: "Table represents the spanner table to write data to.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
	}
}
