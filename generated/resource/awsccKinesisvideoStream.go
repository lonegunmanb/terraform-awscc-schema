package resource

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
        "optional": true,
        "type": "number"
      },
      "device_name": {
        "computed": true,
        "description": "The name of the device that is writing to the stream.",
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
      "kms_key_id": {
        "computed": true,
        "description": "AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "media_type": {
        "computed": true,
        "description": "The media type of the stream. Consumers of the stream can use this information when processing the stream.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Kinesis Video stream.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stream_storage_configuration": {
        "computed": true,
        "description": "Configuration for the storage tier of the Kinesis Video Stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "default_storage_tier": {
              "computed": true,
              "description": "The storage tier for the Kinesis Video Stream. Determines the storage class used for stream data.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. Specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. The following characters can be used: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
    "description": "Resource Type Definition for AWS::KinesisVideo::Stream",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKinesisvideoStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisvideoStream), &result)
	return &result
}
