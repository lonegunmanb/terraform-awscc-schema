package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccS3TablesTable = `{
  "block": {
    "attributes": {
      "compaction": {
        "computed": true,
        "description": "Settings governing the Compaction maintenance action. Contains details about the compaction settings for an Iceberg table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "status": {
              "computed": true,
              "description": "Indicates whether the Compaction maintenance action is enabled.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_file_size_mb": {
              "computed": true,
              "description": "The target file size for the table in MB.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "iceberg_metadata": {
        "computed": true,
        "description": "Contains details about the metadata for an Iceberg table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "iceberg_schema": {
              "computed": true,
              "description": "Contains details about the schema for an Iceberg table",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "schema_field_list": {
                    "computed": true,
                    "description": "Contains details about the schema for an Iceberg table",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "name": {
                          "computed": true,
                          "description": "The name of the field",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "required": {
                          "computed": true,
                          "description": "A Boolean value that specifies whether values are required for each row in this field",
                          "description_kind": "plain",
                          "type": "bool"
                        },
                        "type": {
                          "computed": true,
                          "description": "The field type",
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
      "namespace": {
        "computed": true,
        "description": "The namespace that the table belongs to.",
        "description_kind": "plain",
        "type": "string"
      },
      "open_table_format": {
        "computed": true,
        "description": "Format of the table.",
        "description_kind": "plain",
        "type": "string"
      },
      "snapshot_management": {
        "computed": true,
        "description": "Contains details about the snapshot management settings for an Iceberg table. A snapshot is expired when it exceeds MinSnapshotsToKeep and MaxSnapshotAgeHours.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_snapshot_age_hours": {
              "computed": true,
              "description": "The maximum age of a snapshot before it can be expired.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_snapshots_to_keep": {
              "computed": true,
              "description": "The minimum number of snapshots to keep.",
              "description_kind": "plain",
              "type": "number"
            },
            "status": {
              "computed": true,
              "description": "Indicates whether the SnapshotManagement maintenance action is enabled.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "table_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_bucket_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the specified table bucket.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description": "The name for the table.",
        "description_kind": "plain",
        "type": "string"
      },
      "version_token": {
        "computed": true,
        "description": "The version token of the table",
        "description_kind": "plain",
        "type": "string"
      },
      "warehouse_location": {
        "computed": true,
        "description": "The warehouse location of the table.",
        "description_kind": "plain",
        "type": "string"
      },
      "without_metadata": {
        "computed": true,
        "description": "Indicates that you don't want to specify a schema for the table. This property is mutually exclusive to 'IcebergMetadata', and its only possible value is 'Yes'.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::S3Tables::Table",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccS3TablesTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccS3TablesTable), &result)
	return &result
}
