package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGuarddutyPublishingDestination = `{
  "block": {
    "attributes": {
      "destination_properties": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_arn": {
              "computed": true,
              "description": "The ARN of the resource to publish to.",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_arn": {
              "computed": true,
              "description": "The ARN of the KMS key to use for encryption.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "destination_type": {
        "computed": true,
        "description": "The type of resource for the publishing destination. Currently only Amazon S3 buckets are supported.",
        "description_kind": "plain",
        "type": "string"
      },
      "detector_id": {
        "computed": true,
        "description": "The ID of the GuardDuty detector associated with the publishing destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "publishing_destination_id": {
        "computed": true,
        "description": "The ID of the publishing destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "publishing_failure_start_timestamp": {
        "computed": true,
        "description": "The time, in epoch millisecond format, at which GuardDuty was first unable to publish findings to the destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the publishing destination.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::GuardDuty::PublishingDestination",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGuarddutyPublishingDestinationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGuarddutyPublishingDestination), &result)
	return &result
}
