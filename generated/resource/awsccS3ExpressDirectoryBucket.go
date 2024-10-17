package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3ExpressDirectoryBucket = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Returns the Amazon Resource Name (ARN) of the specified bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "availability_zone_name": {
        "computed": true,
        "description": "Returns the code for the Availability Zone where the directory bucket was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "bucket_encryption": {
        "computed": true,
        "description": "Specifies default encryption for a bucket using server-side encryption with Amazon S3 managed keys (SSE-S3) or AWS KMS keys (SSE-KMS).",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "server_side_encryption_configuration": {
              "computed": true,
              "description": "Specifies the default server-side-encryption configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_key_enabled": {
                    "computed": true,
                    "description": "Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket. Existing objects are not affected. Amazon S3 Express One Zone uses an S3 Bucket Key with SSE-KMS and S3 Bucket Key cannot be disabled. It's only allowed to set the BucketKeyEnabled element to true.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "server_side_encryption_by_default": {
                    "computed": true,
                    "description": "Specifies the default server-side encryption to apply to new objects in the bucket. If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "kms_master_key_id": {
                          "computed": true,
                          "description": "AWS Key Management Service (KMS) customer managed key ID to use for the default encryption. This parameter is allowed only if SSEAlgorithm is set to aws:kms. You can specify this parameter with the key ID or the Amazon Resource Name (ARN) of the KMS key",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "sse_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "bucket_name": {
        "computed": true,
        "description": "Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_redundancy": {
        "description": "Specifies the number of Availability Zone that's used for redundancy for the bucket.",
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
      "location_name": {
        "description": "Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Express::DirectoryBucket.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3ExpressDirectoryBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ExpressDirectoryBucket), &result)
	return &result
}
