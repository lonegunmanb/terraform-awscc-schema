package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKinesisStream = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon resource name (ARN) of the Kinesis stream",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Kinesis stream.",
        "description_kind": "plain",
        "type": "string"
      },
      "retention_period_hours": {
        "computed": true,
        "description": "The number of hours for the data records that are stored in shards to remain accessible.",
        "description_kind": "plain",
        "type": "number"
      },
      "shard_count": {
        "computed": true,
        "description": "The number of shards that the stream uses. Required when StreamMode = PROVISIONED is passed.",
        "description_kind": "plain",
        "type": "number"
      },
      "stream_encryption": {
        "computed": true,
        "description": "When specified, enables or updates server-side encryption using an AWS KMS key for a specified stream.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "The encryption type to use. The only valid value is KMS. ",
              "description_kind": "plain",
              "type": "string"
            },
            "key_id": {
              "computed": true,
              "description": "The GUID for the customer-managed AWS KMS key to use for encryption. This value can be a globally unique identifier, a fully specified Amazon Resource Name (ARN) to either an alias or a key, or an alias name prefixed by \"alias/\".You can also use a master key owned by Kinesis Data Streams by specifying the alias aws/kinesis.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "stream_mode_details": {
        "computed": true,
        "description": "The mode in which the stream is running.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "stream_mode": {
              "computed": true,
              "description": "The mode of the stream",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An arbitrary set of tags (key–value pairs) to associate with the Kinesis stream.",
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
    "description": "Data Source schema for AWS::Kinesis::Stream",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKinesisStreamSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKinesisStream), &result)
	return &result
}
