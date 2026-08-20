package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeChannelFlow = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "description": "The ARN of the app instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_instance_id": {
        "computed": true,
        "description": "The ID of the app instance, extracted from the channel flow ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the channel flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "channel_flow_id": {
        "computed": true,
        "description": "The ID of the channel flow, extracted from the channel flow ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_timestamp": {
        "computed": true,
        "description": "The time at which the channel flow was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The time at which the channel flow was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the channel flow.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "processors": {
        "description": "Information about the processor Lambda functions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration": {
              "description": "A processor's metadata.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda": {
                    "description": "Stores metadata about a Lambda processor.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "invocation_type": {
                          "description": "Controls how the Lambda function is invoked.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "resource_arn": {
                          "description": "The ARN of the Lambda message processing function.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "required": true
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "execution_order": {
              "description": "The sequence in which processors run.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "fallback_action": {
              "description": "Determines whether to continue or stop processing when communication with a processor fails.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description": "The name of the processor.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "The tags for the channel flow.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in a tag.",
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
    "description": "Creates a channel flow in the Amazon Chime SDK Messaging service. A channel flow is a container for processors (Lambda functions) that perform actions on chat messages.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeChannelFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeChannelFlow), &result)
	return &result
}
