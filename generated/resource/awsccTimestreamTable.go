package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTimestreamTable = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "database_name": {
        "description": "The name for the database which the table to be created belongs to.",
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
      "magnetic_store_write_properties": {
        "computed": true,
        "description": "The properties that determine whether magnetic store writes are enabled.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "enable_magnetic_store_writes": {
              "description": "Boolean flag indicating whether magnetic store writes are enabled.",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "magnetic_store_rejected_data_location": {
              "computed": true,
              "description": "Location to store information about records that were asynchronously rejected during magnetic store writes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "s3_configuration": {
                    "computed": true,
                    "description": "S3 configuration for location to store rejections from magnetic store writes",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "bucket_name": {
                          "description": "The bucket name used to store the data.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "encryption_option": {
                          "description": "Either SSE_KMS or SSE_S3.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "kms_key_id": {
                          "computed": true,
                          "description": "Must be provided if SSE_KMS is specified as the encryption option",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "object_key_prefix": {
                          "computed": true,
                          "description": "String used to prefix all data in the bucket.",
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
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "name": {
        "computed": true,
        "description": "The table name exposed as a read-only attribute.",
        "description_kind": "plain",
        "type": "string"
      },
      "retention_properties": {
        "computed": true,
        "description": "The retention duration of the memory store and the magnetic store.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "magnetic_store_retention_period_in_days": {
              "computed": true,
              "description": "The duration for which data must be stored in the magnetic store.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "memory_store_retention_period_in_hours": {
              "computed": true,
              "description": "The duration for which data must be stored in the memory store.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "schema": {
        "computed": true,
        "description": "A Schema specifies the expected data model of the table.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "composite_partition_key": {
              "computed": true,
              "description": "A list of partition keys defining the attributes used to partition the table data. The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "enforcement_in_record": {
                    "computed": true,
                    "description": "The level of enforcement for the specification of a dimension key in ingested records. Options are REQUIRED (dimension key must be specified) and OPTIONAL (dimension key does not have to be specified).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name of the attribute used for a dimension key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "type": {
                    "description": "The type of the partition key. Options are DIMENSION (dimension key) and MEASURE (measure key).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
      "table_name": {
        "computed": true,
        "description": "The name for the table. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the table name.",
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
    "description": "The AWS::Timestream::Table resource creates a Timestream Table.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTimestreamTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTimestreamTable), &result)
	return &result
}
