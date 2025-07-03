package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisStreamConsumer = `{
  "block": {
    "attributes": {
      "consumer_arn": {
        "computed": true,
        "description": "The ARN returned by Kinesis Data Streams when you registered the consumer. If you don't know the ARN of the consumer that you want to deregister, you can use the ListStreamConsumers operation to get a list of the descriptions of all the consumers that are currently registered with a given data stream. The description of a consumer contains its ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "consumer_creation_timestamp": {
        "computed": true,
        "description": "Timestamp when the consumer was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "consumer_name": {
        "computed": true,
        "description": "The name of the Kinesis Stream Consumer. For a given Kinesis data stream, each consumer must have a unique name. However, consumer names don't have to be unique across data streams.",
        "description_kind": "plain",
        "type": "string"
      },
      "consumer_status": {
        "computed": true,
        "description": "A consumer can't read data while in the CREATING or DELETING states. Valid Values: CREATING | DELETING | ACTIVE",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stream_arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the Kinesis data stream that you want to register the consumer with.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key–value pairs) to associate with the Kinesis consumer.",
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
    "description": "Data Source schema for AWS::Kinesis::StreamConsumer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisStreamConsumerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisStreamConsumer), &result)
	return &result
}
