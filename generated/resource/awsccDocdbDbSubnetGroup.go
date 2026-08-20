package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbDbSubnetGroup = `{
  "block": {
    "attributes": {
      "db_subnet_group_description": {
        "description": "The description for the subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "The name for the db subnet group. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "subnet_ids": {
        "description": "One or more subnet IDs to be assigned to the db subnet group.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "One or more tags to be assigned to the db subnet group",
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
    "description": "Resource Type definition for AWS::DocDB::DBSubnetGroup",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDocdbDbSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbDbSubnetGroup), &result)
	return &result
}
