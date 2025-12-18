package resource

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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name for the maintenance window target.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "owner_information": {
        "computed": true,
        "description": "A user-provided value that will be included in any Amazon CloudWatch Events events that are raised while running tasks for these targets in this maintenance window.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_type": {
        "description": "The type of target that is being registered with the maintenance window.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "targets": {
        "description": "The targets to register with the maintenance window.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "User-defined criteria for sending commands that target managed nodes that meet the criteria.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "values": {
              "description": "User-defined criteria that maps to Key.",
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "window_id": {
        "description": "The ID of the maintenance window to register the target with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "window_target_id": {
        "computed": true,
        "description": "The ID of the target.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::SSM::MaintenanceWindowTarget",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmMaintenanceWindowTargetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmMaintenanceWindowTarget), &result)
	return &result
}
