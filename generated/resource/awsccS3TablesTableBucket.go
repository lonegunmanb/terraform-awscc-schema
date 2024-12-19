package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTableBucket = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_name": {
        "description": "A name for the table bucket.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
              "optional": true,
              "type": "number"
            },
            "status": {
              "computed": true,
              "description": "Indicates whether the Unreferenced File Removal maintenance action is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "unreferenced_days": {
              "computed": true,
              "description": "For any object not referenced by your table and older than the UnreferencedDays property, S3 creates a delete marker and marks the object version as noncurrent.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Creates an Amazon S3 Tables table bucket in the same AWS Region where you create the AWS CloudFormation stack.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccS3TablesTableBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTableBucket), &result)
	return &result
}
