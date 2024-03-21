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
        "description_kind": "plain",
        "type": "string"
      },
      "content_based_deduplication": {
        "computed": true,
        "description": "For first-in-first-out (FIFO) queues, specifies whether to enable content-based deduplication. During the deduplication interval, SQS treats messages that are sent with identical content as duplicates and delivers only one copy of the message. For more information, see the ` + "`" + `` + "`" + `ContentBasedDeduplication` + "`" + `` + "`" + ` attribute for the ` + "`" + `` + "`" + `CreateQueue` + "`" + `` + "`" + ` action in the *API Reference*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "deduplication_scope": {
        "computed": true,
        "description": "For high throughput for FIFO queues, specifies whether message deduplication occurs at the message group or queue level. Valid values are ` + "`" + `` + "`" + `messageGroup` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `queue` + "`" + `` + "`" + `.\n To enable high throughput for a FIFO queue, set this attribute to ` + "`" + `` + "`" + `messageGroup` + "`" + `` + "`" + ` *and* set the ` + "`" + `` + "`" + `FifoThroughputLimit` + "`" + `` + "`" + ` attribute to ` + "`" + `` + "`" + `perMessageGroupId` + "`" + `` + "`" + `. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "delay_seconds": {
        "computed": true,
        "description": "The time in seconds for which the delivery of all messages in the queue is delayed. You can specify an integer value of ` + "`" + `` + "`" + `0` + "`" + `` + "`" + ` to ` + "`" + `` + "`" + `900` + "`" + `` + "`" + ` (15 minutes). The default value is ` + "`" + `` + "`" + `0` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "fifo_queue": {
        "computed": true,
        "description": "If set to true, creates a FIFO queue. If you don't specify this property, SQS creates a standard queue. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fifo_throughput_limit": {
        "computed": true,
        "description": "For high throughput for FIFO queues, specifies whether the FIFO queue throughput quota applies to the entire queue or per message group. Valid values are ` + "`" + `` + "`" + `perQueue` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `perMessageGroupId` + "`" + `` + "`" + `.\n To enable high throughput for a FIFO queue, set this attribute to ` + "`" + `` + "`" + `perMessageGroupId` + "`" + `` + "`" + ` *and* set the ` + "`" + `` + "`" + `DeduplicationScope` + "`" + `` + "`" + ` attribute to ` + "`" + `` + "`" + `messageGroup` + "`" + `` + "`" + `. If you set these attributes to anything other than these values, normal throughput is in effect and deduplication occurs as specified. For more information, see [High throughput for FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html) and [Quotas related to messages](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html) in the *Developer Guide*.",
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
        "description": "The length of time in seconds for which SQS can reuse a data key to encrypt or decrypt messages before calling KMS again. The value must be an integer between 60 (1 minute) and 86,400 (24 hours). The default is 300 (5 minutes).\n  A shorter time period provides better security, but results in more calls to KMS, which might incur charges after Free Tier. For more information, see [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kms_master_key_id": {
        "computed": true,
        "description": "The ID of an AWS Key Management Service (KMS) for SQS, or a custom KMS. To use the AWS managed KMS for SQS, specify a (default) alias ARN, alias name (e.g. ` + "`" + `` + "`" + `alias/aws/sqs` + "`" + `` + "`" + `), key ARN, or key ID. For more information, see the following:\n  +   [Encryption at rest](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html) in the *Developer Guide* \n  +   [CreateQueue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/APIReference/API_CreateQueue.html) in the *API Reference* \n  +   [Request Parameters](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters) in the *Key Management Service API Reference* \n  +   The Key Management Service (KMS) section of the [Best Practices](https://docs.aws.amazon.com/https://d0.awsstatic.com/whitepapers/aws-kms-best-practices.pdf) whitepaper",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_message_size": {
        "computed": true,
        "description": "The limit of how many bytes that a message can contain before SQS rejects it. You can specify an integer value from ` + "`" + `` + "`" + `1,024` + "`" + `` + "`" + ` bytes (1 KiB) to ` + "`" + `` + "`" + `262,144` + "`" + `` + "`" + ` bytes (256 KiB). The default value is ` + "`" + `` + "`" + `262,144` + "`" + `` + "`" + ` (256 KiB).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "message_retention_period": {
        "computed": true,
        "description": "The number of seconds that SQS retains a message. You can specify an integer value from ` + "`" + `` + "`" + `60` + "`" + `` + "`" + ` seconds (1 minute) to ` + "`" + `` + "`" + `1,209,600` + "`" + `` + "`" + ` seconds (14 days). The default value is ` + "`" + `` + "`" + `345,600` + "`" + `` + "`" + ` seconds (4 days).",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "queue_name": {
        "computed": true,
        "description": "A name for the queue. To create a FIFO queue, the name of your FIFO queue must end with the ` + "`" + `` + "`" + `.fifo` + "`" + `` + "`" + ` suffix. For more information, see [FIFO queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html) in the *Developer Guide*.\n If you don't specify a name, CFN generates a unique physical ID and uses that ID for the queue name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) in the *User Guide*. \n  If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "queue_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "receive_message_wait_time_seconds": {
        "computed": true,
        "description": "Specifies the duration, in seconds, that the ReceiveMessage action call waits until a message is in the queue in order to include it in the response, rather than returning an empty response if a message isn't yet available. You can specify an integer from 1 to 20. Short polling is used as the default or when you specify 0 for this property. For more information, see [Consuming messages using long polling](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-short-and-long-polling.html#sqs-long-polling) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "redrive_allow_policy": {
        "computed": true,
        "description": "The string that includes the parameters for the permissions for the dead-letter queue redrive permission and which source queues can specify dead-letter queues as a JSON object. The parameters are as follows:\n  +  ` + "`" + `` + "`" + `redrivePermission` + "`" + `` + "`" + `: The permission type that defines which source queues can specify the current queue as the dead-letter queue. Valid values are:\n  +  ` + "`" + `` + "`" + `allowAll` + "`" + `` + "`" + `: (Default) Any source queues in this AWS account in the same Region can specify this queue as the dead-letter queue.\n  +  ` + "`" + `` + "`" + `denyAll` + "`" + `` + "`" + `: No source queues can specify this queue as the dead-letter queue.\n  +  ` + "`" + `` + "`" + `byQueue` + "`" + `` + "`" + `: Only queues specified by the ` + "`" + `` + "`" + `sourceQueueArns` + "`" + `` + "`" + ` parameter can specify this queue as the dead-letter queue.\n  \n  +  ` + "`" + `` + "`" + `sourceQueueArns` + "`" + `` + "`" + `: The Amazon Resource Names (ARN)s of the source queues that can specify this queue as the dead-letter queue and redrive messages. You can specify this parameter only when the ` + "`" + `` + "`" + `redrivePermission` + "`" + `` + "`" + ` parameter is set to ` + "`" + `` + "`" + `byQueue` + "`" + `` + "`" + `. You can specify up to 10 source queue ARNs. To allow more than 10 source queues to specify dead-letter queues, set the ` + "`" + `` + "`" + `redrivePermission` + "`" + `` + "`" + ` parameter to ` + "`" + `` + "`" + `allowAll` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "redrive_policy": {
        "computed": true,
        "description": "The string that includes the parameters for the dead-letter queue functionality of the source queue as a JSON object. The parameters are as follows:\n  +  ` + "`" + `` + "`" + `deadLetterTargetArn` + "`" + `` + "`" + `: The Amazon Resource Name (ARN) of the dead-letter queue to which SQS moves messages after the value of ` + "`" + `` + "`" + `maxReceiveCount` + "`" + `` + "`" + ` is exceeded.\n  +  ` + "`" + `` + "`" + `maxReceiveCount` + "`" + `` + "`" + `: The number of times a message is delivered to the source queue before being moved to the dead-letter queue. When the ` + "`" + `` + "`" + `ReceiveCount` + "`" + `` + "`" + ` for a message exceeds the ` + "`" + `` + "`" + `maxReceiveCount` + "`" + `` + "`" + ` for a queue, SQS moves the message to the dead-letter-queue.\n  \n  The dead-letter queue of a FIFO queue must also be a FIFO queue. Similarly, the dead-letter queue of a standard queue must also be a standard queue.\n   *JSON* \n  ` + "`" + `` + "`" + `{ \"deadLetterTargetArn\" : String, \"maxReceiveCount\" : Integer }` + "`" + `` + "`" + ` \n  *YAML* \n  ` + "`" + `` + "`" + `deadLetterTargetArn : String` + "`" + `` + "`" + ` \n  ` + "`" + `` + "`" + `maxReceiveCount : Integer` + "`" + `` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sqs_managed_sse_enabled": {
        "computed": true,
        "description": "Enables server-side queue encryption using SQS owned encryption keys. Only one server-side encryption option is supported per queue (for example, [SSE-KMS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html) or [SSE-SQS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)). When ` + "`" + `` + "`" + `SqsManagedSseEnabled` + "`" + `` + "`" + ` is not defined, ` + "`" + `` + "`" + `SSE-SQS` + "`" + `` + "`" + ` encryption is enabled by default.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "tags": {
        "computed": true,
        "description": "The tags that you attach to this queue. For more information, see [Resource tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) in the *User Guide*.",
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
        "description": "The length of time during which a message will be unavailable after a message is delivered from the queue. This blocks other components from receiving the same message and gives the initial component time to process and delete the message from the queue.\n Values must be from 0 to 43,200 seconds (12 hours). If you don't specify a value, AWS CloudFormation uses the default value of 30 seconds.\n For more information about SQS queue visibility timeouts, see [Visibility timeout](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html) in the *Developer Guide*.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::SQS::Queue` + "`" + `` + "`" + ` resource creates an SQS standard or FIFO queue.\n Keep the following caveats in mind:\n  +  If you don't specify the ` + "`" + `` + "`" + `FifoQueue` + "`" + `` + "`" + ` property, SQS creates a standard queue.\n  You can't change the queue type after you create it and you can't convert an existing standard queue into a FIFO queue. You must either create a new FIFO queue for your application or delete your existing standard queue and recreate it as a FIFO queue. For more information, see [Moving from a standard queue to a FIFO queue](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-moving.html) in the *Developer Guide*. \n   +  If you don't provide a value for a property, the queue is created with the default value for the property.\n  +  If you delete a queue, you must wait at least 60 seconds before creating a queue with the same name.\n  +  To successfully create a new queue, you must provide a queue name that adheres to the [limits related to queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html) and is unique within the scope of your queues.\n  \n For more information about creating FIFO (first-in-first-out) queues, see [Creating an queue ()](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/screate-queue-cloudformation.html) in the *Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSqsQueueSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSqsQueue), &result)
	return &result
}
