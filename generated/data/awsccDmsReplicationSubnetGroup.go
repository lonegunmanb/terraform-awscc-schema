package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsReplicationSubnetGroup = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_subnet_group_description": {
        "computed": true,
        "description": "The description for the subnet group.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_subnet_group_identifier": {
        "computed": true,
        "description": "The name for the replication subnet group. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "One or more subnet IDs to be assigned to the replication subnet group.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "One or more tags to be assigned to the replication subnet group",
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
    "description": "Data Source schema for AWS::DMS::ReplicationSubnetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDmsReplicationSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsReplicationSubnetGroup), &result)
	return &result
}
