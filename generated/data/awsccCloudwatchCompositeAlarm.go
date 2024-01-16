package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchCompositeAlarm = `{
  "block": {
    "attributes": {
      "actions_enabled": {
        "computed": true,
        "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
        "description_kind": "plain",
        "type": "bool"
      },
      "actions_suppressor": {
        "computed": true,
        "description": "Actions will be suppressed if the suppressor alarm is in the ALARM state. ActionsSuppressor can be an AlarmName or an Amazon Resource Name (ARN) from an existing alarm. ",
        "description_kind": "plain",
        "type": "string"
      },
      "actions_suppressor_extension_period": {
        "computed": true,
        "description": "Actions will be suppressed if WaitPeriod is active. The length of time that actions are suppressed is in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "actions_suppressor_wait_period": {
        "computed": true,
        "description": "Actions will be suppressed if ExtensionPeriod is active. The length of time that actions are suppressed is in seconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "alarm_actions": {
        "computed": true,
        "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "alarm_description": {
        "computed": true,
        "description": "The description of the alarm",
        "description_kind": "plain",
        "type": "string"
      },
      "alarm_name": {
        "computed": true,
        "description": "The name of the Composite Alarm",
        "description_kind": "plain",
        "type": "string"
      },
      "alarm_rule": {
        "computed": true,
        "description": "Expression which aggregates the state of other Alarms (Metric or Composite Alarms)",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the alarm",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "insufficient_data_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ok_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::CompositeAlarm",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchCompositeAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchCompositeAlarm), &result)
	return &result
}
