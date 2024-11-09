package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsDeliveryDestination = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_destination_policy": {
        "computed": true,
        "description": "IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 51200",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_destination_type": {
        "computed": true,
        "description": "Displays whether this delivery destination is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
        "description_kind": "plain",
        "type": "string"
      },
      "destination_resource_arn": {
        "computed": true,
        "description": "The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.",
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
        "description": "The name of this delivery destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "output_format": {
        "computed": true,
        "description": "The format of the logs that are sent to this delivery destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that have been assigned to this delivery destination.",
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
    "description": "Data Source schema for AWS::Logs::DeliveryDestination",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsDeliveryDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDeliveryDestination), &result)
	return &result
}
