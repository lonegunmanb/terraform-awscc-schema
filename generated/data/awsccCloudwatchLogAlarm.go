package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchLogAlarm = `{
  "block": {
    "attributes": {
      "action_log_line_count": {
        "computed": true,
        "description": "The number of log lines to include in alarm notifications. Valid values are 0 to 50.",
        "description_kind": "plain",
        "type": "number"
      },
      "action_log_line_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications. Required when ActionLogLineCount is greater than 0.",
        "description_kind": "plain",
        "type": "string"
      },
      "actions_enabled": {
        "computed": true,
        "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
        "description_kind": "plain",
        "type": "bool"
      },
      "alarm_actions": {
        "computed": true,
        "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "alarm_description": {
        "computed": true,
        "description": "The description of the log alarm.",
        "description_kind": "plain",
        "type": "string"
      },
      "alarm_name": {
        "computed": true,
        "description": "The name of the log alarm.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the log alarm.",
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "computed": true,
        "description": "The arithmetic operation to use when comparing the specified threshold and the query results. Valid values are GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, and LessThanOrEqualToThreshold.",
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
        "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "ok_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the OK state from any other state.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "query_results_to_alarm": {
        "computed": true,
        "description": "The number of query results that must be breaching to trigger the alarm.",
        "description_kind": "plain",
        "type": "number"
      },
      "query_results_to_evaluate": {
        "computed": true,
        "description": "The number of query results over which data is compared to the specified threshold.",
        "description_kind": "plain",
        "type": "number"
      },
      "scheduled_query_configuration": {
        "computed": true,
        "description": "The scheduled query configuration for the log alarm.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aggregation_expression": {
              "computed": true,
              "description": "The aggregation expression for the scheduled query, e.g. count(*) or avg(latency) by host.",
              "description_kind": "plain",
              "type": "string"
            },
            "log_group_identifiers": {
              "computed": true,
              "description": "The log groups to query.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "query_string": {
              "computed": true,
              "description": "The query string to execute against the specified log groups.",
              "description_kind": "plain",
              "type": "string"
            },
            "schedule_configuration": {
              "computed": true,
              "description": "The schedule configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time_offset": {
                    "computed": true,
                    "description": "The number of seconds into the past to end the query window. Must be a non-negative value and cannot exceed 2592000 seconds (30 days).",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "schedule_expression": {
                    "computed": true,
                    "description": "The expression that defines when the scheduled query runs, e.g. rate(1 minute).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time_offset": {
                    "computed": true,
                    "description": "The number of seconds into the past to start the query window. Must be a positive value and cannot exceed 2592000 seconds (30 days).",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "scheduled_query_role_arn": {
              "computed": true,
              "description": "The ARN of the IAM role that grants permissions to execute the scheduled query.",
              "description_kind": "plain",
              "type": "string"
            },
            "tags": {
              "computed": true,
              "description": "A list of key-value pairs to associate with the scheduled query that backs the log alarm.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "A unique identifier for the tag. The combination of tag keys and values can help you organize and categorize your resources.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the specified tag key.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the log alarm.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A unique identifier for the tag. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "threshold": {
        "computed": true,
        "description": "The value to compare against the results of the scheduled query evaluation.",
        "description_kind": "plain",
        "type": "number"
      },
      "treat_missing_data": {
        "computed": true,
        "description": "Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.",
        "description_kind": "plain",
        "type": "string"
      },
      "warm_up_configuration": {
        "computed": true,
        "description": "The warm-up configuration for the alarm. During the warm-up period, the alarm stays in INSUFFICIENT_DATA and doesn't perform alarm actions. For more information, see Alarm warm-up periods in the Amazon CloudWatch User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "only_start_evaluating_after_warm_up_period_ends": {
              "computed": true,
              "description": "Specifies whether the alarm waits for the full warm-up period before it starts evaluating. If true, the alarm waits the entire WarmUpPeriodDurationInMinutes before it starts evaluating, even if metric data arrives earlier. If false, the alarm ends the warm-up period early and starts evaluating as soon as it has enough metric data to fill its evaluation window. This is the default behavior.",
              "description_kind": "plain",
              "type": "bool"
            },
            "warm_up_period_duration_in_minutes": {
              "computed": true,
              "description": "The length of the warm-up period, in minutes. For this duration after you create or update the alarm, the alarm stays in INSUFFICIENT_DATA and doesn't perform alarm actions. Valid values range from 1 to 2,880 minutes (2 days). You can change this value while the alarm is still in its warm-up period. Changes have no effect after the warm-up period ends.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::LogAlarm",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchLogAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchLogAlarm), &result)
	return &result
}
