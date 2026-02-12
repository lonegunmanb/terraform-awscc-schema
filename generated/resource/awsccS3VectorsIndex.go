package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3VectorsIndex = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description": "Date and time when the vector index was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_type": {
        "description": "The data type of the vectors to be inserted into the vector index.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dimension": {
        "description": "The dimensions of the vectors to be inserted into the vector index.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "distance_metric": {
        "description": "The distance metric to be used for similarity search.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_configuration": {
        "computed": true,
        "description": "The encryption configuration for the index.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "AWS Key Management Service (KMS) customer managed key ID to use for the encryption configuration. This parameter is allowed if and only if sseType is set to aws:kms",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sse_type": {
              "computed": true,
              "description": "Defines the server-side encryption type for index encryption configuration. Defaults to the parent vector bucket's encryption settings when unspecified.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "index_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the index",
        "description_kind": "plain",
        "type": "string"
      },
      "index_name": {
        "computed": true,
        "description": "The name of the vector index to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metadata_configuration": {
        "computed": true,
        "description": "The metadata configuration for the vector index.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "non_filterable_metadata_keys": {
              "computed": true,
              "description": "Non-filterable metadata keys allow you to enrich vectors with additional context during storage and retrieval. Unlike default metadata keys, these keys cannot be used as query filters. Non-filterable metadata keys can be retrieved but cannot be searched, queried, or filtered. You can access non-filterable metadata keys of your vectors after finding the vectors.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "User tags (key-value pairs) to associate with the index.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key must be between 1 to 128 characters in length. Tag key cannot start with 'aws:' and can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value must be between 0 to 256 characters in length. Tag value can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "vector_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the vector bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vector_bucket_name": {
        "computed": true,
        "description": "The name of the vector bucket that contains the vector index.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::S3Vectors::Index",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3VectorsIndexSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3VectorsIndex), &result)
	return &result
}
