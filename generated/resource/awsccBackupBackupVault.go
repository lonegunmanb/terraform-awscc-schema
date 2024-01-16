package resource

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
        "optional": true,
        "type": "string"
      },
      "backup_vault_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "backup_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "backup_vault_tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "encryption_key_arn": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
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
              "optional": true,
              "type": "number"
            },
            "max_retention_days": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_retention_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "notifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "backup_vault_events": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sns_topic_arn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Backup::BackupVault",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupBackupVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupBackupVault), &result)
	return &result
}
