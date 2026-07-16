package resource

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
        "optional": true,
        "type": "number"
      },
      "action_log_line_role_arn": {
        "computed": true,
        "description": "The ARN of the IAM role that grants CloudWatch permissions to fetch log lines for alarm notifications. Required when ActionLogLineCount is greater than 0.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "actions_enabled": {
        "computed": true,
        "description": "Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "alarm_actions": {
        "computed": true,
        "description": "The list of actions to execute when this alarm transitions into an ALARM state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "alarm_description": {
        "computed": true,
        "description": "The description of the log alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "alarm_name": {
        "computed": true,
        "description": "The name of the log alarm.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the log alarm.",
        "description_kind": "plain",
        "type": "string"
      },
      "comparison_operator": {
        "description": "The arithmetic operation to use when comparing the specified threshold and the query results. Valid values are GreaterThanOrEqualToThreshold, GreaterThanThreshold, LessThanThreshold, and LessThanOrEqualToThreshold.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "insufficient_data_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "ok_actions": {
        "computed": true,
        "description": "The actions to execute when this alarm transitions to the OK state from any other state.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "query_results_to_alarm": {
        "description": "The number of query results that must be breaching to trigger the alarm.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "query_results_to_evaluate": {
        "description": "The number of query results over which data is compared to the specified threshold.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "scheduled_query_configuration": {
        "description": "The scheduled query configuration for the log alarm.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "aggregation_expression": {
              "description": "The aggregation expression for the scheduled query, e.g. count(*) or avg(latency) by host.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "log_group_identifiers": {
              "computed": true,
              "description": "The log groups to query.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_string": {
              "description": "The query string to execute against the specified log groups.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "schedule_configuration": {
              "description": "The schedule configuration.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time_offset": {
                    "computed": true,
                    "description": "The number of seconds into the past to end the query window. Must be a non-negative value and cannot exceed 2592000 seconds (30 days).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "schedule_expression": {
                    "description": "The expression that defines when the scheduled query runs, e.g. rate(1 minute).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time_offset": {
                    "description": "The number of seconds into the past to start the query window. Must be a positive value and cannot exceed 2592000 seconds (30 days).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "scheduled_query_role_arn": {
              "description": "The ARN of the IAM role that grants permissions to execute the scheduled query.",
              "description_kind": "plain",
              "required": true,
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
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the specified tag key.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "threshold": {
        "description": "The value to compare against the results of the scheduled query evaluation.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "treat_missing_data": {
        "computed": true,
        "description": "Sets how this alarm is to handle missing data points. Valid values are breaching, notBreaching, ignore, and missing.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::LogAlarm. A LogAlarm evaluates scheduled query results from CloudWatch Logs and triggers actions when thresholds are breached.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchLogAlarmSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchLogAlarm), &result)
	return &result
}
