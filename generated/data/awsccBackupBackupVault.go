package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupBackupVault = `{
  "block": {
    "attributes": {
      "access_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "encryption_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "lock_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "changeable_for_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "max_retention_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "min_retention_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "notifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "backup_vault_events": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "sns_topic_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Backup::BackupVault",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupBackupVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupBackupVault), &result)
	return &result
}
