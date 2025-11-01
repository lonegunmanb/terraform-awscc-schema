package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2CapacityManagerDataExport = `{
  "block": {
    "attributes": {
      "capacity_manager_data_export_id": {
        "computed": true,
        "description": "The unique identifier of the capacity manager data export.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "output_format": {
        "description": "The format of the exported capacity manager data. Choose 'csv' for comma-separated values or 'parquet' for optimized columnar storage format.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket_name": {
        "description": "The name of the Amazon S3 bucket where the capacity manager data export will be stored. The bucket must exist and be accessible by EC2 Capacity Manager service.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_bucket_prefix": {
        "computed": true,
        "description": "The prefix for the S3 bucket location where exported files will be placed. If not specified, files will be placed in the root of the bucket.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "schedule": {
        "description": "The schedule for the capacity manager data export. Currently supports hourly exports that provide periodic snapshots of capacity manager data.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the capacity manager data export.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::EC2::CapacityManagerDataExport",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccEc2CapacityManagerDataExportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2CapacityManagerDataExport), &result)
	return &result
}
