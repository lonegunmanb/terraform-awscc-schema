package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSmsvoiceConfigurationSet = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "configuration_set_name": {
        "computed": true,
        "description": "The name to use for the configuration set.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_sender_id": {
        "computed": true,
        "description": "The default sender ID to set for the ConfigurationSet.",
        "description_kind": "plain",
        "type": "string"
      },
      "event_destinations": {
        "computed": true,
        "description": "An event destination is a location where you send message events.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "cloudwatch_logs_destination": {
              "computed": true,
              "description": "An object that contains IamRoleArn and LogGroupArn associated with an Amazon CloudWatch event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iam_role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of an AWS Identity and Access Management role that is able to write event data to an Amazon CloudWatch destination.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "log_group_arn": {
                    "computed": true,
                    "description": "The name of the Amazon CloudWatch log group that you want to record events in.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "enabled": {
              "computed": true,
              "description": "When set to true events will be logged. By default this is set to true",
              "description_kind": "plain",
              "type": "bool"
            },
            "event_destination_name": {
              "computed": true,
              "description": "The name that identifies the event destination.",
              "description_kind": "plain",
              "type": "string"
            },
            "kinesis_firehose_destination": {
              "computed": true,
              "description": "An object that contains IamRoleArn and DeliveryStreamArn associated with an Amazon Kinesis Firehose event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "delivery_stream_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the delivery stream.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "iam_role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of an AWS Identity and Access Management role that is able to write event data to an Amazon CloudWatch destination.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "matching_event_types": {
              "computed": true,
              "description": "An array of event types that determine which events to log. If 'ALL' is used, then AWS End User Messaging SMS and Voice logs every event type.",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            },
            "sns_destination": {
              "computed": true,
              "description": "An object that contains SNS TopicArn event destination.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "topic_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish events to.",
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
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "message_feedback_enabled": {
        "computed": true,
        "description": "Set to true to enable message feedback.",
        "description_kind": "plain",
        "type": "bool"
      },
      "protect_configuration_id": {
        "computed": true,
        "description": "The unique identifier for the protect configuration to be associated to the configuration set.",
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
    "description": "Data Source schema for AWS::SMSVOICE::ConfigurationSet",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSmsvoiceConfigurationSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSmsvoiceConfigurationSet), &result)
	return &result
}
