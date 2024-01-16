package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectRoutingProfile = `{
  "block": {
    "attributes": {
      "agent_availability_timer": {
        "computed": true,
        "description": "Whether agents with this routing profile will have their routing order calculated based on longest idle time or time since their last inbound contact.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_outbound_queue_arn": {
        "computed": true,
        "description": "The identifier of the default outbound queue for this routing profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the routing profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_concurrencies": {
        "computed": true,
        "description": "The channels agents can handle in the Contact Control Panel (CCP) for this routing profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "channel": {
              "computed": true,
              "description": "The channels that agents can handle in the Contact Control Panel (CCP).",
              "description_kind": "plain",
              "type": "string"
            },
            "concurrency": {
              "computed": true,
              "description": "The number of contacts an agent can have on a channel simultaneously.",
              "description_kind": "plain",
              "type": "number"
            },
            "cross_channel_behavior": {
              "computed": true,
              "description": "Defines the cross-channel routing behavior that allows an agent working on a contact in one channel to be offered a contact from a different channel.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "behavior_type": {
                    "computed": true,
                    "description": "Specifies the other channels that can be routed to an agent handling their current channel.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "name": {
        "computed": true,
        "description": "The name of the routing profile.",
        "description_kind": "plain",
        "type": "string"
      },
      "queue_configs": {
        "computed": true,
        "description": "The queues to associate with this routing profile.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "delay": {
              "computed": true,
              "description": "The delay, in seconds, a contact should wait in the queue before they are routed to an available agent.",
              "description_kind": "plain",
              "type": "number"
            },
            "priority": {
              "computed": true,
              "description": "The order in which contacts are to be handled for the queue.",
              "description_kind": "plain",
              "type": "number"
            },
            "queue_reference": {
              "computed": true,
              "description": "Contains the channel and queue identifier for a routing profile.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "channel": {
                    "computed": true,
                    "description": "The channels that agents can handle in the Contact Control Panel (CCP).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "queue_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) for the queue.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "list"
        }
      },
      "routing_profile_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the routing profile.",
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
      }
    },
    "description": "Data Source schema for AWS::Connect::RoutingProfile",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectRoutingProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectRoutingProfile), &result)
	return &result
}
