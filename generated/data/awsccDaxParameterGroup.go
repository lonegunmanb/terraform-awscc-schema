package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDaxParameterGroup = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description of the parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameter_group_name": {
        "computed": true,
        "description": "The name of the parameter group.",
        "description_kind": "plain",
        "type": "string"
      },
      "parameter_name_values": {
        "computed": true,
        "description": "An array of name-value pairs for the parameters in the group. Each element in the array represents a single parameter.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DAX::ParameterGroup",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDaxParameterGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDaxParameterGroup), &result)
	return &result
}
