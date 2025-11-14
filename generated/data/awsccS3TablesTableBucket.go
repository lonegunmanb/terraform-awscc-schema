package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTableBucket = `{
  "block": {
    "attributes": {
      "encryption_configuration": {
        "computed": true,
        "description": "Specifies encryption settings for the table bucket",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "kms_key_arn": {
              "computed": true,
              "description": "ARN of the KMS key to use for encryption",
              "description_kind": "plain",
              "type": "string"
            },
            "sse_algorithm": {
              "computed": true,
              "description": "Server-side encryption algorithm",
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
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_name": {
        "computed": true,
        "description": "A name for the table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "User tags (key-value pairs) to associate with the table bucket.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "Tag key must be between 1 to 128 characters in length. Tag key cannot start with 'aws:' and can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "Tag value must be between 0 to 256 characters in length. Tag value can only contain alphanumeric characters, spaces, _, ., /, =, +, -, and @.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "unreferenced_file_removal": {
        "computed": true,
        "description": "Settings governing the Unreferenced File Removal maintenance action. Unreferenced file removal identifies and deletes all objects that are not referenced by any table snapshots.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "noncurrent_days": {
              "computed": true,
              "description": "S3 permanently deletes noncurrent objects after the number of days specified by the NoncurrentDays property.",
              "description_kind": "plain",
              "type": "number"
            },
            "status": {
              "computed": true,
              "description": "Indicates whether the Unreferenced File Removal maintenance action is enabled.",
              "description_kind": "plain",
              "type": "string"
            },
            "unreferenced_days": {
              "computed": true,
              "description": "For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::S3Tables::TableBucket",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3TablesTableBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTableBucket), &result)
	return &result
}
