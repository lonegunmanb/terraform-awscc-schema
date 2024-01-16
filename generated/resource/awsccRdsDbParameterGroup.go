package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccRdsDbParameterGroup = `{
  "block": {
    "attributes": {
      "db_parameter_group_name": {
        "computed": true,
        "description": "Specifies the name of the DB parameter group",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description": "Provides the customer-specified description for this DB parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "family": {
        "description": "The DB parameter group family name.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "An array of parameter names and values for the parameter update.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "The AWS::RDS::DBParameterGroup resource creates a custom parameter group for an RDS database family",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccRdsDbParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccRdsDbParameterGroup), &result)
	return &result
}
