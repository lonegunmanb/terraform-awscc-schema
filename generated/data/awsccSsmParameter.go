package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmParameter = `{
  "block": {
    "attributes": {
      "allowed_pattern": {
        "computed": true,
        "description": "The regular expression used to validate the parameter value.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_type": {
        "computed": true,
        "description": "The corresponding DataType of the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The information about the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "The policies attached to the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "computed": true,
        "description": "The corresponding tier of the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of the parameter.",
        "description_kind": "plain",
        "type": "string"
      },
      "value": {
        "computed": true,
        "description": "The value associated with the parameter.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::Parameter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmParameterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmParameter), &result)
	return &result
}
