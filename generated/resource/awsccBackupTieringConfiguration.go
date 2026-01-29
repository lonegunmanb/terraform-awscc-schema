package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupTieringConfiguration = `{
  "block": {
    "attributes": {
      "backup_vault_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "creation_time": {
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
      "last_updated_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_selection": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "resource_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resources": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "tiering_down_settings_in_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "tiering_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tiering_configuration_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tiering_configuration_tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::Backup::TieringConfiguration",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupTieringConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupTieringConfiguration), &result)
	return &result
}
