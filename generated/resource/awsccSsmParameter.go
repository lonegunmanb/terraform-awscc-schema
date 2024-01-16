package resource

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
        "optional": true,
        "type": "string"
      },
      "data_type": {
        "computed": true,
        "description": "The corresponding DataType of the parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The information about the parameter.",
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
      "name": {
        "computed": true,
        "description": "The name of the parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policies": {
        "computed": true,
        "description": "The policies attached to the parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A key-value pair to associate with a resource.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tier": {
        "computed": true,
        "description": "The corresponding tier of the parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "type": {
        "description": "The type of the parameter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "value": {
        "description": "The value associated with the parameter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SSM::Parameter",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmParameterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmParameter), &result)
	return &result
}
