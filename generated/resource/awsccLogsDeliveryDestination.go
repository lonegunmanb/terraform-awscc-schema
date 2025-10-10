package resource

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
        "nested_type": {
          "attributes": {
            "delivery_destination_name": {
              "computed": true,
              "description": "The name of the delivery destination to assign this policy to",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delivery_destination_policy": {
              "computed": true,
              "description": "The contents of the policy attached to the delivery destination",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "delivery_destination_type": {
        "computed": true,
        "description": "Displays whether this delivery destination is CloudWatch Logs, Amazon S3, Kinesis Data Firehose, or XRay.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "destination_resource_arn": {
        "computed": true,
        "description": "The ARN of the Amazon Web Services destination that this delivery destination represents. That Amazon Web Services destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a delivery stream in Firehose.",
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
      "name": {
        "description": "The name of this delivery destination.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "output_format": {
        "computed": true,
        "description": "The format of the logs that are sent to this delivery destination.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "This structure contains information about one delivery destination in your account.\n\nA delivery destination is an AWS resource that represents an AWS service that logs can be sent to CloudWatch Logs, Amazon S3, are supported as Kinesis Data Firehose delivery destinations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsDeliveryDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDeliveryDestination), &result)
	return &result
}
