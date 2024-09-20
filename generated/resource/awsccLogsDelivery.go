package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsDelivery = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_destination_arn": {
        "description": "The ARN of the delivery destination that is associated with this delivery.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delivery_destination_type": {
        "computed": true,
        "description": "Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, or Kinesis Data Firehose.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_id": {
        "computed": true,
        "description": "The unique ID that identifies this delivery in your account.",
        "description_kind": "plain",
        "type": "string"
      },
      "delivery_source_name": {
        "description": "The name of the delivery source that is associated with this delivery.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags that have been assigned to this delivery.",
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
    "description": "This structure contains information about one delivery in your account.\n\nA delivery is a connection between a logical delivery source and a logical delivery destination.\n\nFor more information, see [CreateDelivery](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_CreateDelivery.html).",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsDeliverySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsDelivery), &result)
	return &result
}
