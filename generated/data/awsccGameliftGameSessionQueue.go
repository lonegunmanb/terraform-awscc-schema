package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftGameSessionQueue = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift game session queue resource and uniquely identifies it.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_event_data": {
        "computed": true,
        "description": "Information that is added to all events that are related to this game session queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "destinations": {
        "computed": true,
        "description": "A list of fleets and/or fleet aliases that can be used to fulfill game session placement requests in the queue.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "filter_configuration": {
        "computed": true,
        "description": "A list of locations where a queue is allowed to place new game sessions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allowed_locations": {
              "computed": true,
              "description": "A list of locations to allow game session placement in, in the form of AWS Region codes such as us-west-2.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A descriptive label that is associated with game session queue. Queue names must be unique within each Region.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_target": {
        "computed": true,
        "description": "An SNS topic ARN that is set up to receive game session placement notifications.",
        "description_kind": "plain",
        "type": "string"
      },
      "player_latency_policies": {
        "computed": true,
        "description": "A set of policies that act as a sliding cap on player latency.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "maximum_individual_player_latency_milliseconds": {
              "computed": true,
              "description": "The maximum latency value that is allowed for any player, in milliseconds. All policies must have a value set for this property.",
              "description_kind": "plain",
              "type": "number"
            },
            "policy_duration_seconds": {
              "computed": true,
              "description": "The length of time, in seconds, that the policy is enforced while placing a new game session.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "priority_configuration": {
        "computed": true,
        "description": "Custom settings to use when prioritizing destinations and locations for game session placements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "location_order": {
              "computed": true,
              "description": "The prioritization order to use for fleet locations, when the PriorityOrder property includes LOCATION.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "priority_order": {
              "computed": true,
              "description": "The recommended sequence to use when prioritizing where to place new game sessions.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
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
      },
      "timeout_in_seconds": {
        "computed": true,
        "description": "The maximum time, in seconds, that a new game session placement request remains in the queue.",
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::GameLift::GameSessionQueue",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGameliftGameSessionQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftGameSessionQueue), &result)
	return &result
}
