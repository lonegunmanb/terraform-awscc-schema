package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisvideoStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Kinesis Video stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_retention_in_hours": {
        "computed": true,
        "description": "The number of hours till which Kinesis Video will retain the data in the stream",
        "description_kind": "plain",
        "type": "number"
      },
      "device_name": {
        "computed": true,
        "description": "The name of the device that is writing to the stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_id": {
        "computed": true,
        "description": "AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.",
        "description_kind": "plain",
        "type": "string"
      },
      "media_type": {
        "computed": true,
        "description": "The media type of the stream. Consumers of the stream can use this information when processing the stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Kinesis Video stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs associated with the Kinesis Video Stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. Specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::KinesisVideo::Stream",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisvideoStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisvideoStream), &result)
	return &result
}
