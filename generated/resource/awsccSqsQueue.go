package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSqsQueue = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_based_deduplication": {
        "computed": true,
        "description": "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, Amazon SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deduplication_scope": {
        "computed": true,
        "description": "Specifies whether message deduplication occurs at the message group or queue level. Valid values are messageGroup and queue.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delay_seconds": {
        "computed": true,
        "description": "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of 0 to 900 (15 minutes). The default value is 0.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fifo_queue": {
        "computed": true,
        "description": "If set to true, creates a FIFO queue. If you don't specify this property, Amazon SQS creates a standard queue.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fifo_throughput_limit": {
        "computed": true,
        "description": "Specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are perQueue and perMessageGroupId. The perMessageGroupId value is allowed only when the value for DeduplicationScope is messageGroup.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "kms_data_key_reuse_period_seconds": {
        "computed": true,
        "description": "The length of time in seconds for which Amazon SQS can reuse a data key to encrypt or decrypt messages before calling AWS KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_master_key_id": {
        "computed": true,
        "description": "The ID of an AWS managed customer master key (CMK) for Amazon SQS or a custom CMK. To use the AWS managed CMK for Amazon SQS, specify the (default) alias alias/aws/sqs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_message_size": {
        "computed": true,
        "description": "The limit of how many bytes that a message can contain before Amazon SQS rejects it. You can specify an integer value from 1,024 bytes (1 KiB) to 262,144 bytes (256 KiB). The default value is 262,144 (256 KiB).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_retention_period": {
        "computed": true,
        "description": "The number of seconds that Amazon SQS retains a message. You can specify an integer value from 60 seconds (1 minute) to 1,209,600 seconds (14 days). The default value is 345,600 seconds (4 days).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queue_name": {
        "computed": true,
        "description": "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the .fifo suffix.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queue_url": {
        "computed": true,
        "description": "URL of the source queue.",
        "description_kind": "plain",
        "type": "string"
      },
      "receive_message_wait_time_seconds": {
        "computed": true,
        "description": "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "redrive_allow_policy": {
        "computed": true,
        "description": "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redrive_policy": {
        "computed": true,
        "description": "A string that includes the parameters for the dead-letter queue functionality (redrive policy) of the source queue.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sqs_managed_sse_enabled": {
        "computed": true,
        "description": "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (e.g. SSE-KMS or SSE-SQS ).",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this queue.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "visibility_timeout": {
        "computed": true,
        "description": "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue. Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::SQS::Queue",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSqsQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSqsQueue), &result)
	return &result
}
