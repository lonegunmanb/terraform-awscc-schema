package resource

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
              "optional": true,
              "type": "string"
            },
            "provisioned_throughput": {
              "computed": true,
              "description": "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "read_capacity_units": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "write_capacity_units": {
                    "description_kind": "plain",
                    "required": true,
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
      "client_side_timestamps_enabled": {
        "computed": true,
        "description": "Indicates whether client side timestamps are enabled (true) or disabled (false) on the table. False by default, once it is enabled it cannot be disabled again.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "clustering_key_columns": {
        "computed": true,
        "description": "Clustering key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column": {
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "column_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "column_type": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "order_by": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "default_time_to_live": {
        "computed": true,
        "description": "Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "string"
            },
            "kms_key_identifier": {
              "computed": true,
              "description": "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
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
      "keyspace_name": {
        "description": "Name for Cassandra keyspace",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "partition_key_columns": {
        "description": "Partition key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "column_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "point_in_time_recovery_enabled": {
        "computed": true,
        "description": "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "regular_columns": {
        "computed": true,
        "description": "Non-key columns of the table",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "column_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "column_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "table_name": {
        "computed": true,
        "description": "Name for Cassandra table",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
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
    "description": "Resource schema for AWS::Cassandra::Table",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCassandraTableSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCassandraTable), &result)
	return &result
}
