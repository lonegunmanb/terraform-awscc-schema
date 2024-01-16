package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A unique identifier for the matchmaking rule set.",
        "description_kind": "plain",
        "type": "string"
      },
      "rule_set_body": {
        "computed": true,
        "description": "A collection of matchmaking rules, formatted as a JSON string.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::GameLift::MatchmakingRuleSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGameliftMatchmakingRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftMatchmakingRuleSet), &result)
	return &result
}
