package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCassandraKeyspace = `{
  "block": {
    "attributes": {
      "client_side_timestamps_enabled": {
        "computed": true,
        "description": "Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace. To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can?t disable it again.",
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keyspace_name": {
        "computed": true,
        "description": "Name for Cassandra keyspace",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_specification": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "region_list": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "replication_strategy": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Cassandra::Keyspace",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCassandraKeyspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCassandraKeyspace), &result)
	return &result
}
