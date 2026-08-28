package resource

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
        "description": "The name to be assigned to the adapter being created.",
        "description_kind": "plain",
        "required": true,
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
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "feature_types": {
        "description": "The type of feature that the adapter is being trained on. Currently, supported feature types are: QUERIES",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
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
    "description": "The AWS::Textract::Adapter resource creates an Amazon Textract adapter, which can be fine-tuned for enhanced performance on user-provided documents.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTextractAdapterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTextractAdapter), &result)
	return &result
}
