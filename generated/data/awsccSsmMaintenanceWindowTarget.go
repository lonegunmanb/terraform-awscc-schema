package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmMaintenanceWindowTarget = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description for the target.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the maintenance window target.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner_information": {
        "computed": true,
        "description": "A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_type": {
        "computed": true,
        "description": "The type of target that is being registered with the maintenance window.",
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description": "The targets to register with the maintenance window.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "User-defined criteria for sending commands that target managed nodes that meet the criteria.",
              "description_kind": "plain",
              "type": "string"
            },
            "values": {
              "computed": true,
              "description": "User-defined criteria that maps to Key.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "window_id": {
        "computed": true,
        "description": "The ID of the maintenance window to register the target with.",
        "description_kind": "plain",
        "type": "string"
      },
      "window_target_id": {
        "computed": true,
        "description": "The ID of the target.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::MaintenanceWindowTarget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmMaintenanceWindowTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmMaintenanceWindowTarget), &result)
	return &result
}
