package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeChannelFlow = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "computed": true,
        "description": "The ARN of the app instance.",
        "description_kind": "plain",
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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The time at which the channel flow was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the channel flow.",
        "description_kind": "plain",
        "type": "string"
      },
      "processors": {
        "computed": true,
        "description": "Information about the processor Lambda functions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "configuration": {
              "computed": true,
              "description": "A processor's metadata.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "lambda": {
                    "computed": true,
                    "description": "Stores metadata about a Lambda processor.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "invocation_type": {
                          "computed": true,
                          "description": "Controls how the Lambda function is invoked.",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "resource_arn": {
                          "computed": true,
                          "description": "The ARN of the Lambda message processing function.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "single"
              }
            },
            "execution_order": {
              "computed": true,
              "description": "The sequence in which processors run.",
              "description_kind": "plain",
              "type": "number"
            },
            "fallback_action": {
              "computed": true,
              "description": "Determines whether to continue or stop processing when communication with a processor fails.",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the processor.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in a tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Chime::ChannelFlow",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChimeChannelFlowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeChannelFlow), &result)
	return &result
}
