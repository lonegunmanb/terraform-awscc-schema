package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGlueDatabase = `{
  "block": {
    "attributes": {
      "catalog_id": {
        "computed": true,
        "description": "The AWS account ID for the account in which to create the catalog object.",
        "description_kind": "plain",
        "type": "string"
      },
      "database_input": {
        "computed": true,
        "description": "The metadata for the database.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "create_table_default_permissions": {
              "computed": true,
              "description": "Creates a set of default permissions on the table for principals. Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "permissions": {
                    "computed": true,
                    "description": "The permissions that are granted to the principal.",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "principal": {
                    "computed": true,
                    "description": "The principal who is granted permissions.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "data_lake_principal_identifier": {
                          "computed": true,
                          "description": "An identifier for the AWS Lake Formation principal.",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    }
                  }
                },
                "nesting_mode": "list"
              }
            },
            "description": {
              "computed": true,
              "description": "A description of the database.",
              "description_kind": "plain",
              "type": "string"
            },
            "federated_database": {
              "computed": true,
              "description": "A FederatedDatabase structure that references an entity outside the AWS Glue Data Catalog.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "connection_name": {
                    "computed": true,
                    "description": "The name of the connection to the external metastore.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "identifier": {
                    "computed": true,
                    "description": "A unique identifier for the federated database.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "location_uri": {
              "computed": true,
              "description": "The location of the database (for example, an HDFS path).",
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "computed": true,
              "description": "The name of the database. For hive compatibility, this is folded to lowercase when it is stored.",
              "description_kind": "plain",
              "type": "string"
            },
            "parameters": {
              "computed": true,
              "description": "These key-value pairs define parameters and properties of the database.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_database": {
              "computed": true,
              "description": "A DatabaseIdentifier structure that describes a target database for resource linking.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "catalog_id": {
                    "computed": true,
                    "description": "The ID of the Data Catalog in which the database resides.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "database_name": {
                    "computed": true,
                    "description": "The name of the catalog database.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "region": {
                    "computed": true,
                    "description": "Region of the target database.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "database_name": {
        "computed": true,
        "description": "The name of the database. For hive compatibility, this is folded to lowercase when it is store.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Glue::Database",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGlueDatabaseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGlueDatabase), &result)
	return &result
}
