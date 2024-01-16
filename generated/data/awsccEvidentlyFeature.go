package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEvidentlyFeature = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_variation": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "entity_overrides": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "entity_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "variation": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "evaluation_strategy": {
        "computed": true,
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
        "description_kind": "plain",
        "type": "string"
      },
      "project": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "variations": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "boolean_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "double_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "long_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "string_value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "variation_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Evidently::Feature",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEvidentlyFeatureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEvidentlyFeature), &result)
	return &result
}
