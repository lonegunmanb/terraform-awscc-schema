package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCassandraTable = `{
  "block": {
    "attributes": {
      "billing_mode": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "mode": {
              "computed": true,
              "description": "Capacity mode for the specified table",
              "description_kind": "plain",
              "type": "string"
            },
            "provisioned_throughput": {
              "computed": true,
              "description": "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "write_capacity_units": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "client_side_timestamps_enabled": {
        "computed": true,
        "description": "Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.",
        "description_kind": "plain",
        "type": "bool"
      },
      "clustering_key_columns": {
        "computed": true,
        "description": "Clustering key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "column_type": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "order_by": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "default_time_to_live": {
        "computed": true,
        "description": "Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.",
        "description_kind": "plain",
        "type": "number"
      },
      "encryption_specification": {
        "computed": true,
        "description": "Represents the settings used to enable server-side encryption",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "encryption_type": {
              "computed": true,
              "description": "Server-side encryption type",
              "description_kind": "plain",
              "type": "string"
            },
            "kms_key_identifier": {
              "computed": true,
              "description": "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
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
      "keyspace_name": {
        "computed": true,
        "description": "Name for Cassandra keyspace",
        "description_kind": "plain",
        "type": "string"
      },
      "partition_key_columns": {
        "computed": true,
        "description": "Partition key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "column_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "point_in_time_recovery_enabled": {
        "computed": true,
        "description": "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
        "description_kind": "plain",
        "type": "bool"
      },
      "regular_columns": {
        "computed": true,
        "description": "Non-key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "column_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "table_name": {
        "computed": true,
        "description": "Name for Cassandra table",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Cassandra::Table",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCassandraTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCassandraTable), &result)
	return &result
}