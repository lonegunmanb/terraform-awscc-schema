package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupRestoreTestingSelection = `{
  "block": {
    "attributes": {
      "iam_role_arn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "protected_resource_arns": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
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
            "string_not_equals": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "protected_resource_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_metadata_overrides": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "restore_testing_plan_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_testing_selection_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "validation_window_hours": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "description": "Resource Type definition for AWS::Backup::RestoreTestingSelection",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBackupRestoreTestingSelectionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupRestoreTestingSelection), &result)
	return &result
}
