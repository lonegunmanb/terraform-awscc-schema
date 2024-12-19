package data

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
        "description": "Returns the code for the Availability Zone or Local Zone where the directory bucket was created. An example for the code of an Availability Zone is 'us-east-1f'.",
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
                          "type": "string"
                        },
                        "sse_algorithm": {
                          "computed": true,
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "bucket_name": {
        "computed": true,
        "description": "Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone or Local Zone. The bucket name must also follow the format 'bucket_base_name--zone_id--x-s3'. The zone_id can be the ID of an Availability Zone or a Local Zone. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_redundancy": {
        "computed": true,
        "description": "Specifies the number of Availability Zone or Local Zone that's used for redundancy for the bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lifecycle_configuration": {
        "computed": true,
        "description": "Lifecycle rules that define how Amazon S3 Express manages objects during their lifetime.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "rules": {
              "computed": true,
              "description": "A lifecycle rule for individual objects in an Amazon S3 Express bucket.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "abort_incomplete_multipart_upload": {
                    "computed": true,
                    "description": "Specifies the days since the initiation of an incomplete multipart upload that Amazon S3 will wait before permanently removing all parts of the upload.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "days_after_initiation": {
                          "computed": true,
                          "description": "Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.",
                          "description_kind": "plain",
                          "type": "number"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  },
                  "expiration_in_days": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_size_greater_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_size_less_than": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "prefix": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "location_name": {
        "computed": true,
        "description": "Specifies the Zone ID of the Availability Zone or Local Zone where the directory bucket will be created. An example Availability Zone ID value is 'use1-az5'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Express::DirectoryBucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3ExpressDirectoryBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3ExpressDirectoryBucket), &result)
	return &result
}
