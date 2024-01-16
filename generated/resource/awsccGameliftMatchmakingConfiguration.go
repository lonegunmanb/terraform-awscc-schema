package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftMatchmakingConfiguration = `{
  "block": {
    "attributes": {
      "acceptance_required": {
        "description": "A flag that indicates whether a match that was created with this configuration must be accepted by the matched players",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "acceptance_timeout_seconds": {
        "computed": true,
        "description": "The length of time (in seconds) to wait for players to accept a proposed match, if acceptance is required.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "additional_player_count": {
        "computed": true,
        "description": "The number of player slots in a match to keep open for future players.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift matchmaking configuration resource and uniquely identifies it.",
        "description_kind": "plain",
        "type": "string"
      },
      "backfill_mode": {
        "computed": true,
        "description": "The method used to backfill game sessions created with this matchmaking configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description": "A time stamp indicating when this data object was created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_event_data": {
        "computed": true,
        "description": "Information to attach to all events related to the matchmaking configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A descriptive label that is associated with matchmaking configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "flex_match_mode": {
        "computed": true,
        "description": "Indicates whether this matchmaking configuration is being used with Amazon GameLift hosting or as a standalone matchmaking solution.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "game_properties": {
        "computed": true,
        "description": "A set of custom properties for a game session, formatted as key:value pairs.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The game property identifier.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The game property value.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "game_session_data": {
        "computed": true,
        "description": "A set of custom game session properties, formatted as a single string value.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "game_session_queue_arns": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.",
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
      "name": {
        "description": "A unique identifier for the matchmaking configuration.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_target": {
        "computed": true,
        "description": "An SNS topic ARN that is set up to receive matchmaking notifications.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "request_timeout_seconds": {
        "description": "The maximum duration, in seconds, that a matchmaking ticket can remain in process before timing out.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "rule_set_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) associated with the GameLift matchmaking rule set resource that this configuration uses.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule_set_name": {
        "description": "A unique identifier for the matchmaking rule set to use with this configuration.",
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
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
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
    "description": "The AWS::GameLift::MatchmakingConfiguration resource creates an Amazon GameLift (GameLift) matchmaking configuration.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftMatchmakingConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftMatchmakingConfiguration), &result)
	return &result
}
