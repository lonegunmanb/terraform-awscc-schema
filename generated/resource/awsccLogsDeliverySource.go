package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsDeliverySource = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery source.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "log_type": {
        "computed": true,
        "description": "The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The unique name of the Log source.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "The ARN of the resource that will be sending the logs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_arns": {
        "computed": true,
        "description": "This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "service": {
        "computed": true,
        "description": "The AWS service that is sending logs.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that have been assigned to this delivery source.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode",
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
    "description": " A delivery source is an AWS resource that sends logs to an AWS destination. The destination can be CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.\n\nOnly some AWS services support being configured as a delivery source. These services are listed as Supported [V2 Permissions] in the table at [Enabling logging from AWS services](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-and-resource-policy.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsDeliverySourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDeliverySource), &result)
	return &result
}
