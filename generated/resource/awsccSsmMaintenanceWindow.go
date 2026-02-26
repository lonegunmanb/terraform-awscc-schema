package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmMaintenanceWindow = `{
  "block": {
    "attributes": {
      "allow_unassociated_targets": {
        "description": "Enables a maintenance window task to run on managed instances, even if you have not registered those instances as targets. If enabled, then you must specify the unregistered instances (by instance ID) when you register a task with the maintenance window.",
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "cutoff": {
        "description": "The number of hours before the end of the maintenance window that AWS Systems Manager stops scheduling new tasks for execution.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "description": {
        "computed": true,
        "description": "A description of the maintenance window.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "duration": {
        "description": "The duration of the maintenance window in hours.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "end_date": {
        "computed": true,
        "description": "The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become inactive.",
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
        "description": "The name of the maintenance window.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "description": "The schedule of the maintenance window in the form of a cron or rate expression.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_offset": {
        "computed": true,
        "description": "The number of days to wait to run a maintenance window after the scheduled cron expression date and time.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "schedule_timezone": {
        "computed": true,
        "description": "The time zone that the scheduled maintenance window executions are based on, in Internet Assigned Numbers Authority (IANA) format.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "start_date": {
        "computed": true,
        "description": "The date and time, in ISO-8601 Extended format, for when the maintenance window is scheduled to become active. StartDate allows you to delay activation of the maintenance window until the specified future date.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Optional metadata that you assign to a resource in the form of an arbitrary set of tags (key-value pairs). Tags enable you to categorize a resource in different ways, such as by purpose, owner, or environment. For example, you might want to tag a maintenance window to identify the type of tasks it will run, the types of targets, and the environment it will run in.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "window_id": {
        "computed": true,
        "description": "The ID of the maintenance window.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::SSM::MaintenanceWindow",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmMaintenanceWindowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmMaintenanceWindow), &result)
	return &result
}
