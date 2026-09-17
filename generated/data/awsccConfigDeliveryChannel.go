package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConfigDeliveryChannel = `{
  "block": {
    "attributes": {
      "config_snapshot_delivery_properties": {
        "computed": true,
        "description": "The options for how often AWS Config delivers configuration snapshots to the Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "delivery_frequency": {
              "computed": true,
              "description": "The frequency with which AWS Config delivers configuration snapshots.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the delivery channel. By default, AWS Config assigns the name \"default\" when creating the delivery channel. To change the delivery channel name, you must use the DeleteDeliveryChannel action to delete your current delivery channel, and then you must use the PutDeliveryChannel command to create a delivery channel that has the desired name.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_bucket_name": {
        "computed": true,
        "description": "The name of the Amazon S3 bucket to which AWS Config delivers configuration snapshots and configuration history files.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_key_prefix": {
        "computed": true,
        "description": "The prefix for the specified Amazon S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_kms_key_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS ) AWS KMS key (KMS key) used to encrypt objects delivered by AWS Config. Must belong to the same Region as the destination S3 bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "sns_topic_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon SNS topic to which AWS Config sends notifications about configuration changes.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Config::DeliveryChannel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConfigDeliveryChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConfigDeliveryChannel), &result)
	return &result
}
