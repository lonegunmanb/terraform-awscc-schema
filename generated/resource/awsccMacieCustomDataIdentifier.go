package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Custom data identifier ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "ignore_words": {
        "computed": true,
        "description": "Words to be ignored.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "keywords": {
        "computed": true,
        "description": "Keywords to be matched against.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "maximum_match_distance": {
        "computed": true,
        "description": "Maximum match distance.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Name of custom data identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "regex": {
        "description": "Regular expression for custom data identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
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
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Macie CustomDataIdentifier resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMacieCustomDataIdentifierSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieCustomDataIdentifier), &result)
	return &result
}
