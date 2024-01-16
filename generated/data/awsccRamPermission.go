package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRamPermission = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_resource_type_default": {
        "computed": true,
        "description": "Set to true to use this as the default permission.",
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "computed": true,
        "description": "The name of the permission.",
        "description_kind": "plain",
        "type": "string"
      },
      "permission_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "policy_template": {
        "computed": true,
        "description": "Policy template for the permission.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The resource type this permission can be used with.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "version": {
        "computed": true,
        "description": "Version of the permission.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::RAM::Permission",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccRamPermissionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRamPermission), &result)
	return &result
}
