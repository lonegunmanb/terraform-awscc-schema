package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupLogicallyAirGappedBackupVault = `{
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
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "max_retention_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "min_retention_days": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "notifications": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "backup_vault_events": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "sns_topic_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "vault_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vault_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Backup::LogicallyAirGappedBackupVault",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupLogicallyAirGappedBackupVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupLogicallyAirGappedBackupVault), &result)
	return &result
}
