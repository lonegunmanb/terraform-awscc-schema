package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieCustomDataIdentifier = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Custom data identifier ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "Description of custom data identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ignore_words": {
        "computed": true,
        "description": "Words to be ignored.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "keywords": {
        "computed": true,
        "description": "Keywords to be matched against.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "maximum_match_distance": {
        "computed": true,
        "description": "Maximum match distance.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "Name of custom data identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "regex": {
        "computed": true,
        "description": "Regular expression for custom data identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
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
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Macie::CustomDataIdentifier",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMacieCustomDataIdentifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieCustomDataIdentifier), &result)
	return &result
}
