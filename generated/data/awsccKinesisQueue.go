package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisQueue = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "queue_arn": {
        "computed": true,
        "description": "The ARN of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "queue_attributes": {
        "computed": true,
        "description": "Mutable attributes that control queue consumption behavior.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "dead_letter_configuration": {
              "computed": true,
              "description": "The dead-letter configuration applied to records that exceed the maximum receive count.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "max_receive_count": {
                    "computed": true,
                    "description": "The maximum number of times a record can be received before it is moved to the dead-letter target. Required when TargetType is BUILT_IN.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "target_type": {
                    "computed": true,
                    "description": "The dead-letter target type. BUILT_IN uses a service-managed dead-letter target; NONE disables dead-lettering.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "visibility_timeout_seconds": {
              "computed": true,
              "description": "The duration, in seconds, that a received record is hidden from subsequent receive calls after being handed out to a consumer. Defaults to 30 seconds.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "queue_creation_timestamp": {
        "computed": true,
        "description": "Timestamp of when the queue was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "queue_name": {
        "computed": true,
        "description": "The name of the queue. The name's uniqueness is scoped per AWS account and region.",
        "description_kind": "plain",
        "type": "string"
      },
      "queue_status": {
        "computed": true,
        "description": "The status of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "stream_configuration_list": {
        "computed": true,
        "description": "List of stream configurations associated with the queue. v1 supports a single element; the list shape allows future extensibility to fan in from multiple streams.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "starting_position": {
              "computed": true,
              "description": "The position in the stream from which the queue starts reading records.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "timestamp": {
                    "computed": true,
                    "description": "The timestamp to start reading from. Required when Type is AT_TIMESTAMP; the timestamp must not be in the future.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "type": {
                    "computed": true,
                    "description": "The starting position type.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "stream_arn": {
              "computed": true,
              "description": "The Amazon resource name (ARN) of the Kinesis data stream that the queue reads from.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key-value pairs) to associate with the Kinesis queue.",
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
              "description": "The value for the tag. You can specify a value that is 0 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Kinesis::Queue",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisQueue), &result)
	return &result
}
