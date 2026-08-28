package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTextractAdapter = `{
  "block": {
    "attributes": {
      "adapter_id": {
        "computed": true,
        "description": "A unique identifier for the adapter resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "adapter_name": {
        "computed": true,
        "description": "The name to be assigned to the adapter being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the adapter.",
        "description_kind": "plain",
        "type": "string"
      },
      "auto_update": {
        "computed": true,
        "description": "Controls whether or not the adapter should automatically update.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "The date and time that the adapter was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description to be assigned to the adapter being created.",
        "description_kind": "plain",
        "type": "string"
      },
      "feature_types": {
        "computed": true,
        "description": "The type of feature that the adapter is being trained on. Currently, supported feature types are: QUERIES",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags to be added to the adapter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Textract::Adapter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTextractAdapterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTextractAdapter), &result)
	return &result
}
