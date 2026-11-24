package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAthenaSystemTables = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The time the configuration was created, set by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "kms_key_arn": {
        "computed": true,
        "description": "The ARN of the customer managed AWS KMS key used to encrypt the system table data. Specify a key ARN (not an alias); the key must be in the same account and Region. The encryption choice is permanent for the account and Region: it cannot be changed or removed after creation, and it persists after this resource is deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_at": {
        "computed": true,
        "description": "The time the configuration was last modified, set by the service.",
        "description_kind": "plain",
        "type": "string"
      },
      "table_name": {
        "computed": true,
        "description": "The reserved keyword identifying the account-level System Tables configuration. The only accepted value is ALL, which enables all system tables for the account. This is the resource's primary identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to the system tables configuration.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Athena::SystemTables",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAthenaSystemTablesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAthenaSystemTables), &result)
	return &result
}
