package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMemorydbParameterGroup = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the parameter group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "family": {
        "description": "The name of the parameter group family that this parameter group is compatible with.",
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
      "parameter_group_name": {
        "description": "The name of the parameter group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "An map of parameter names and values for the parameter update. You must supply at least one parameter name and value; subsequent arguments are optional.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this parameter group.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key for the tag. May not be null.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value. May be null.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::MemoryDB::ParameterGroup resource creates an Amazon MemoryDB ParameterGroup.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMemorydbParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMemorydbParameterGroup), &result)
	return &result
}
