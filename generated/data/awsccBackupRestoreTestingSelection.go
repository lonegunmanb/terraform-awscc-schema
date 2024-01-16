package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupRestoreTestingSelection = `{
  "block": {
    "attributes": {
      "iam_role_arn": {
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
      "protected_resource_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "protected_resource_conditions": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "string_equals": {
              "computed": true,
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
            },
            "string_not_equals": {
              "computed": true,
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
          "nesting_mode": "single"
        }
      },
      "protected_resource_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "restore_metadata_overrides": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "restore_testing_plan_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "restore_testing_selection_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validation_window_hours": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "description": "Data Source schema for AWS::Backup::RestoreTestingSelection",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupRestoreTestingSelectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupRestoreTestingSelection), &result)
	return &result
}
