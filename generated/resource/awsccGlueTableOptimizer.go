package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueTableOptimizer = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "description": "The catalog ID of the table",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "database_name": {
        "description": "The name of the database. For Hive compatibility, this is folded to lowercase when it is stored.",
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
      "table_name": {
        "description": "The table name. For Hive compatibility, this must be entirely lowercase.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "table_optimizer_configuration": {
        "description": "Specifies configuration details of a table optimizer.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "compaction_configuration": {
              "computed": true,
              "description": "The configuration for a compaction optimizer. This configuration defines how data files in your table will be compacted to improve query performance and reduce storage costs.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iceberg_configuration": {
                    "computed": true,
                    "description": "The configuration for an Iceberg compaction optimizer.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "delete_file_threshold": {
                          "computed": true,
                          "description": "The minimum number of deletes in a data file to make it eligible for compaction.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_input_files": {
                          "computed": true,
                          "description": "The minimum number of input files before compaction is triggered.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "strategy": {
                          "computed": true,
                          "description": "The compaction strategy to use. Valid values are binpack, sort, and z-order.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "enabled": {
              "description": "Whether the table optimization is enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "orphan_file_deletion_configuration": {
              "computed": true,
              "description": "OrphanFileDeletionConfiguration is a property that can be included within the TableOptimizer resource. It controls the automatic deletion of orphaned files - files that are not tracked by the table metadata, and older than the configured age limit.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iceberg_configuration": {
                    "computed": true,
                    "description": "The IcebergConfiguration property helps optimize your Iceberg tables in AWS Glue by allowing you to specify format-specific settings that control how data is stored, compressed, and managed.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "location": {
                          "computed": true,
                          "description": "Specifies a directory in which to look for orphan files (defaults to the table's location). You may choose a sub-directory rather than the top-level table location.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "orphan_file_retention_period_in_days": {
                          "computed": true,
                          "description": "The specific number of days you want to keep the orphan files.",
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "retention_configuration": {
              "computed": true,
              "description": "The configuration for a snapshot retention optimizer for Apache Iceberg tables.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "iceberg_configuration": {
                    "computed": true,
                    "description": "The configuration for an Iceberg snapshot retention optimizer.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "clean_expired_files": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "number_of_snapshots_to_retain": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "snapshot_retention_period_in_days": {
                          "computed": true,
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
                "nesting_mode": "single"
              },
              "optional": true
            },
            "role_arn": {
              "description": "A role passed by the caller which gives the service permission to update the resources associated with the optimizer on the caller's behalf.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "vpc_configuration": {
              "computed": true,
              "description": "An object that describes the VPC configuration for a table optimizer. This configuration is necessary to perform optimization on tables that are in a customer VPC.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "glue_connection_name": {
                    "computed": true,
                    "description": "The name of the AWS Glue connection used for the VPC for the table optimizer.",
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
          "nesting_mode": "single"
        },
        "required": true
      },
      "type": {
        "description": "The type of table optimizer.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Glue::TableOptimizer",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGlueTableOptimizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueTableOptimizer), &result)
	return &result
}
