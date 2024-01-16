package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAutoscalingScheduledAction = `{
  "block": {
    "attributes": {
      "auto_scaling_group_name": {
        "description": "The name of the Auto Scaling group.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "desired_capacity": {
        "computed": true,
        "description": "The desired capacity is the initial capacity of the Auto Scaling group after the scheduled action runs and the capacity it attempts to maintain.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "end_time": {
        "computed": true,
        "description": "The latest scheduled start time to return. If scheduled action names are provided, this parameter is ignored.",
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
      "max_size": {
        "computed": true,
        "description": "The minimum size of the Auto Scaling group.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "min_size": {
        "computed": true,
        "description": "The minimum size of the Auto Scaling group.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "recurrence": {
        "computed": true,
        "description": "The recurring schedule for the action, in Unix cron syntax format. When StartTime and EndTime are specified with Recurrence , they form the boundaries of when the recurring action starts and stops.",
        "description_kind": "plain",
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description": "The time zone for the cron expression.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::AutoScaling::ScheduledAction resource specifies an Amazon EC2 Auto Scaling scheduled action so that the Auto Scaling group can change the number of instances available for your application in response to predictable load changes.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAutoscalingScheduledActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAutoscalingScheduledAction), &result)
	return &result
}
