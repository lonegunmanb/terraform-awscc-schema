package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotDimension = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN (Amazon resource name) of the created dimension.",
        "description_kind": "plain",
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
        "description": "A unique identifier for the dimension.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "string_values": {
        "description": "Specifies the value or list of values for the dimension.",
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Metadata that can be used to manage the dimension.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag's key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag's value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "type": {
        "description": "Specifies the type of the dimension.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "A dimension can be used to limit the scope of a metric used in a security profile for AWS IoT Device Defender.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIotDimensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotDimension), &result)
	return &result
}
