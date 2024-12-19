package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCassandraType = `{
  "block": {
    "attributes": {
      "direct_parent_types": {
        "computed": true,
        "description": "List of parent User-Defined Types that directly reference the User-Defined Type in their fields.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "direct_referring_tables": {
        "computed": true,
        "description": "List of Tables that directly reference the User-Defined Type in their columns.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "fields": {
        "computed": true,
        "description": "Field definitions of the User-Defined Type",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "field_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "field_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "keyspace_arn": {
        "computed": true,
        "description": "ARN of the Keyspace which contains the User-Defined Type.",
        "description_kind": "plain",
        "type": "string"
      },
      "keyspace_name": {
        "computed": true,
        "description": "Name of the Keyspace which contains the User-Defined Type.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_timestamp": {
        "computed": true,
        "description": "Timestamp of the last time the User-Defined Type's meta data was modified.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_nesting_depth": {
        "computed": true,
        "description": "Maximum nesting depth of the User-Defined Type across the field types.",
        "description_kind": "plain",
        "type": "number"
      },
      "type_name": {
        "computed": true,
        "description": "Name of the User-Defined Type.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cassandra::Type",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCassandraTypeSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCassandraType), &result)
	return &result
}
