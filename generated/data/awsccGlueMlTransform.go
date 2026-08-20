package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueMlTransform = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A user-defined, long-form description text for the machine learning transform.",
        "description_kind": "plain",
        "type": "string"
      },
      "glue_version": {
        "computed": true,
        "description": "The version of AWS Glue this machine learning transform is compatible with.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "input_record_tables": {
        "computed": true,
        "description": "A list of AWS Glue table definitions used by the transform.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "glue_tables": {
              "computed": true,
              "description": "The database and table in the AWS Glue Data Catalog that is used for input or output data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "A unique identifier for the AWS Glue Data Catalog.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to the AWS Glue Data Catalog.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "A database name in the AWS Glue Data Catalog.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "table_name": {
                    "computed": true,
                    "description": "A table name in the AWS Glue Data Catalog.",
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
      },
      "max_capacity": {
        "computed": true,
        "description": "The number of AWS Glue DPUs allocated to task runs for this transform.",
        "description_kind": "plain",
        "type": "number"
      },
      "max_retries": {
        "computed": true,
        "description": "The maximum number of times to retry after an MLTaskRun fails.",
        "description_kind": "plain",
        "type": "number"
      },
      "name": {
        "computed": true,
        "description": "A user-defined name for the machine learning transform.",
        "description_kind": "plain",
        "type": "string"
      },
      "number_of_workers": {
        "computed": true,
        "description": "The number of workers of a defined workerType that are allocated when a task runs.",
        "description_kind": "plain",
        "type": "number"
      },
      "role": {
        "computed": true,
        "description": "The name or ARN of the IAM role with the required permissions.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to use with this machine learning transform.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "timeout": {
        "computed": true,
        "description": "The timeout in minutes of the machine learning transform.",
        "description_kind": "plain",
        "type": "number"
      },
      "transform_encryption": {
        "computed": true,
        "description": "The encryption-at-rest settings of the transform.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "ml_user_data_encryption": {
              "computed": true,
              "description": "The encryption-at-rest settings of the transform that apply to accessing user data.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "kms_key_id": {
                    "computed": true,
                    "description": "The ID for the customer-provided KMS key.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "ml_user_data_encryption_mode": {
                    "computed": true,
                    "description": "The encryption mode applied to user data.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "task_run_security_configuration_name": {
              "computed": true,
              "description": "The name of the security configuration.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "transform_id": {
        "computed": true,
        "description": "The unique identifier for the transform.",
        "description_kind": "plain",
        "type": "string"
      },
      "transform_parameters": {
        "computed": true,
        "description": "The algorithm-specific parameters that are associated with the machine learning transform.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "find_matches_parameters": {
              "computed": true,
              "description": "The parameters to configure the find matches transform.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "accuracy_cost_tradeoff": {
                    "computed": true,
                    "description": "The value for accuracy and cost tradeoff. A value of 0.5 means balance.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "enforce_provided_labels": {
                    "computed": true,
                    "description": "If true, forces the output to match the provided labels.",
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "precision_recall_tradeoff": {
                    "computed": true,
                    "description": "The value for precision and recall tradeoff. A value of 0.5 means no preference.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "primary_key_column_name": {
                    "computed": true,
                    "description": "The name of a column that uniquely identifies rows in the source table.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "transform_type": {
              "computed": true,
              "description": "The type of machine learning transform.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "worker_type": {
        "computed": true,
        "description": "The type of predefined worker that is allocated when a task runs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::MLTransform",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueMlTransformSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueMlTransform), &result)
	return &result
}
