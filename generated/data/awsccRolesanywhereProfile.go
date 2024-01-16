package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRolesanywhereProfile = `{
  "block": {
    "attributes": {
      "duration_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_policy_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "require_instance_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "role_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "session_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
    "description": "Data Source schema for AWS::RolesAnywhere::Profile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRolesanywhereProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRolesanywhereProfile), &result)
	return &result
}
