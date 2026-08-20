package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDocdbDbSubnetGroup = `{
  "block": {
    "attributes": {
      "db_subnet_group_description": {
        "computed": true,
        "description": "The description for the subnet group.",
        "description_kind": "plain",
        "type": "string"
      },
      "db_subnet_group_name": {
        "computed": true,
        "description": "The name for the db subnet group. This value is stored as a lowercase string.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_ids": {
        "computed": true,
        "description": "One or more subnet IDs to be assigned to the db subnet group.",
        "description_kind": "plain",
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
    "description": "Data Source schema for AWS::DocDB::DBSubnetGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDocdbDbSubnetGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDocdbDbSubnetGroup), &result)
	return &result
}
