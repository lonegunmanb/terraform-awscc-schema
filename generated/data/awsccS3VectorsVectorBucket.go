package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3VectorsVectorBucket = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Date and time when the vector bucket was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the vector bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration. This parameter is allowed if and only if sseType is set to aws:kms",
              "description_kind": "plain",
              "type": "string"
            },
            "sse_type": {
              "computed": true,
              "description": "The server-side encryption type to use for the encryption configuration of the vector bucket. By default, if you don't specify, all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3 managed keys (SSE-S3), specifically AES256.",
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
      "vector_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the vector bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "vector_bucket_name": {
        "computed": true,
        "description": "The name of the vector bucket.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Vectors::VectorBucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3VectorsVectorBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3VectorsVectorBucket), &result)
	return &result
}
