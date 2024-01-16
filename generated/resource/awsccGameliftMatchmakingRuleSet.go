package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftMatchmakingRuleSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking rule set resource and uniquely identifies it.",
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "A time stamp indicating when this data object was created. Format is a number expressed in Unix time as milliseconds.",
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
        "description": "A unique identifier for the matchmaking rule set.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rule_set_body": {
        "description": "A collection of matchmaking rules, formatted as a JSON string.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "The AWS::GameLift::MatchmakingRuleSet resource creates an Amazon GameLift (GameLift) matchmaking rule set.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftMatchmakingRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftMatchmakingRuleSet), &result)
	return &result
}
