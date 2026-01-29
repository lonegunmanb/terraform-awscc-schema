package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupTieringConfiguration = `{
  "block": {
    "attributes": {
      "backup_vault_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
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
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_selection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "resources": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "tiering_down_settings_in_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tiering_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tiering_configuration_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tiering_configuration_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Backup::TieringConfiguration",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupTieringConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupTieringConfiguration), &result)
	return &result
}
