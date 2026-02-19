package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchAlarmMuteRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Amazon Resource Name (ARN) of the AlarmMuteRule",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the AlarmMuteRule",
        "description_kind": "plain",
        "type": "string"
      },
      "expire_date": {
        "computed": true,
        "description": "The date, with the same timezone offset as \"ScheduleTimezone\" after which the alarm mute rule will be expired.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The last update timestamp of the alarm mute schedule",
        "description_kind": "plain",
        "type": "string"
      },
      "mute_targets": {
        "computed": true,
        "description": "Targets to be muted",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "alarm_names": {
              "computed": true,
              "description": "The alarm names to be mute by the AlarmMuteRule",
              "description_kind": "plain",
              "type": [
                "set",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "mute_type": {
        "computed": true,
        "description": "The mute type of the alarm mute ",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the AlarmMuteRule",
        "description_kind": "plain",
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description": "The rule for the mute",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule": {
              "computed": true,
              "description": "Schedule for the mute to be active",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "duration": {
                    "computed": true,
                    "description": "The duration of the schedule when it triggers",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expression": {
                    "computed": true,
                    "description": "The expression of the schedule",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "timezone": {
                    "computed": true,
                    "description": "The timezone of the schedule",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "start_date": {
        "computed": true,
        "description": "The date, with the same timezone offset as \"ScheduleTimezone\", after which the alarm mute rule will become active.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the AlarmMuteRule",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
    "description": "Data Source schema for AWS::CloudWatch::AlarmMuteRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchAlarmMuteRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchAlarmMuteRule), &result)
	return &result
}
