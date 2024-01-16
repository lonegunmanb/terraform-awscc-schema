package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatasyncLocationS3 = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket location.",
        "description_kind": "plain",
        "type": "string"
      },
      "location_uri": {
        "computed": true,
        "description": "The URL of the S3 location that was described.",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the Amazon S3 bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "s3_config": {
        "description": "The Amazon Resource Name (ARN) of the AWS IAM role that is used to access an Amazon S3 bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket_access_role_arn": {
              "description": "The ARN of the IAM role of the Amazon S3 bucket.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "s3_storage_class": {
        "computed": true,
        "description": "The Amazon S3 storage class you want to store your files in when this location is used as a task destination.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subdirectory": {
        "computed": true,
        "description": "A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is used to read data from the S3 source location or write data to the S3 destination.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key for an AWS resource tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for an AWS resource tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource schema for AWS::DataSync::LocationS3",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatasyncLocationS3Schema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatasyncLocationS3), &result)
	return &result
}
