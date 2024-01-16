package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBackupRestoreTestingPlan = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recovery_point_selection": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "algorithm": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "exclude_vaults": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "include_vaults": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "recovery_point_types": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "selection_window_days": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "restore_testing_plan_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "restore_testing_plan_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "schedule_expression_timezone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "start_window_hours": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Backup::RestoreTestingPlan",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccBackupRestoreTestingPlanSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBackupRestoreTestingPlan), &result)
	return &result
}
