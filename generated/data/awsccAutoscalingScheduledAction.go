package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingScheduledAction = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "computed": true,
        "description": "The name of the Auto Scaling group.",
        "description_kind": "plain",
        "type": "string"
      },
      "desired_capacity": {
        "computed": true,
        "description": "The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.",
        "description_kind": "plain",
        "type": "number"
      },
      "end_time": {
        "computed": true,
        "description": "The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_size": {
        "computed": true,
        "description": "The minimum size of the Auto Scaling group.",
        "description_kind": "plain",
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description": "The minimum size of the Auto Scaling group.",
        "description_kind": "plain",
        "type": "number"
      },
      "recurrence": {
        "computed": true,
        "description": "The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.",
        "description_kind": "plain",
        "type": "string"
      },
      "scheduled_action_name": {
        "computed": true,
        "description": "Auto-generated unique identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "start_time": {
        "computed": true,
        "description": "The earliest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.",
        "description_kind": "plain",
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description": "The time zone for the cron expression.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AutoScaling::ScheduledAction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAutoscalingScheduledActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingScheduledAction), &result)
	return &result
}
