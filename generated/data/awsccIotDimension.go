package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for the dimension.",
        "description_kind": "plain",
        "type": "string"
      },
      "string_values": {
        "computed": true,
        "description": "Specifies the value or list of values for the dimension.",
        "description_kind": "plain",
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
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "type": {
        "computed": true,
        "description": "Specifies the type of the dimension.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IoT::Dimension",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotDimensionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotDimension), &result)
	return &result
}
