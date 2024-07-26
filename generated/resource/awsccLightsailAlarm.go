package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLightsailAlarm = `{
  "block": {
    "attributes": {
      "alarm_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "alarm_name": {
        "description": "The name for the alarm. Specify the name of an existing alarm to update, and overwrite the previous configuration of the alarm.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "comparison_operator": {
        "description": "The arithmetic operation to use when comparing the specified statistic to the threshold. The specified statistic value is used as the first operand.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "contact_protocols": {
        "computed": true,
        "description": "The contact protocols to use for the alarm, such as Email, SMS (text messaging), or both.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "datapoints_to_alarm": {
        "computed": true,
        "description": "The number of data points that must be not within the specified threshold to trigger the alarm. If you are setting an \"M out of N\" alarm, this value (datapointsToAlarm) is the M.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "evaluation_periods": {
        "description": "The number of most recent periods over which data is compared to the specified threshold. If you are setting an \"M out of N\" alarm, this value (evaluationPeriods) is the N.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "metric_name": {
        "description": "The name of the metric to associate with the alarm.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitored_resource_name": {
        "description": "The name of the Lightsail resource that the alarm monitors.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_enabled": {
        "computed": true,
        "description": "Indicates whether the alarm is enabled. Notifications are enabled by default if you don't specify this parameter.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "notification_triggers": {
        "computed": true,
        "description": "The alarm states that trigger a notification.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "state": {
        "computed": true,
        "description": "The current state of the alarm.",
        "description_kind": "plain",
        "type": "string"
      },
      "threshold": {
        "description": "The value against which the specified statistic is compared.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "treat_missing_data": {
        "computed": true,
        "description": "Sets how this alarm will handle missing data points.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lightsail::Alarm",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLightsailAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLightsailAlarm), &result)
	return &result
}
