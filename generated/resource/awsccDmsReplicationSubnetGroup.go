package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDmsReplicationSubnetGroup = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "replication_subnet_group_description": {
        "description": "The description for the subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replication_subnet_group_identifier": {
        "computed": true,
        "description": "The name for the replication subnet group. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_ids": {
        "description": "One or more subnet IDs to be assigned to the replication subnet group.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::DMS::ReplicationSubnetGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDmsReplicationSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDmsReplicationSubnetGroup), &result)
	return &result
}
