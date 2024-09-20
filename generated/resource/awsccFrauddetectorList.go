package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccFrauddetectorList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The list ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description": "The time when the list was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the list.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "elements": {
        "computed": true,
        "description": "The elements in this list.",
        "description_kind": "plain",
        "optional": true,
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
      "last_updated_time": {
        "computed": true,
        "description": "The time when the list was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the list.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with this list.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
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
      },
      "variable_type": {
        "computed": true,
        "description": "The variable type of the list.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "A resource schema for a List in Amazon Fraud Detector.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccFrauddetectorListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccFrauddetectorList), &result)
	return &result
}
