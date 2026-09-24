package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBcmScheduledReport = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the scheduled report.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The timestamp when the scheduled report was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_arn": {
        "description": "The ARN of the dashboard associated with the scheduled report. Managed dashboards cannot be used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description of the scheduled report's purpose or contents.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "health_status": {
        "computed": true,
        "description": "The health status of the scheduled report at its last refresh time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "last_refreshed_at": {
              "computed": true,
              "description": "The time at which the health status was last refreshed.",
              "description_kind": "plain",
              "type": "string"
            },
            "status_code": {
              "computed": true,
              "description": "Whether the scheduled report is healthy.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the scheduled report.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule_config": {
        "description": "The schedule configuration that defines when and how often the report is generated.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule_expression": {
              "computed": true,
              "description": "The schedule expression that specifies when to trigger the scheduled report run. This value must be a cron expression consisting of six fields separated by white spaces: cron(minutes hours day_of_month month day_of_week year).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_expression_time_zone": {
              "computed": true,
              "description": "The time zone for the schedule expression, for example, UTC.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_period": {
              "computed": true,
              "description": "The time period during which the schedule is active.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time": {
                    "computed": true,
                    "description": "The time at which the schedule stops being active.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "start_time": {
                    "computed": true,
                    "description": "The time at which the schedule becomes active.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "state": {
              "computed": true,
              "description": "The state of the schedule. ENABLED means the scheduled report runs according to its schedule expression. DISABLED means the scheduled report is paused and will not run until re-enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "scheduled_report_execution_role_arn": {
        "description": "The ARN of the IAM role that the scheduled report uses to execute. AWS Billing and Cost Management Dashboards assumes this IAM role while executing the scheduled report.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags applied to the scheduled report.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag value.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the scheduled report was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "widget_date_range_override": {
        "computed": true,
        "description": "The date range override applied to widgets in the scheduled report.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_time": {
              "computed": true,
              "description": "The end of the range.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "Whether Value is an absolute date or a duration relative to now.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The date, or an ISO 8601 duration when Type is RELATIVE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "start_time": {
              "computed": true,
              "description": "The start of the range.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "type": {
                    "computed": true,
                    "description": "Whether Value is an absolute date or a duration relative to now.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The date, or an ISO 8601 duration when Type is RELATIVE.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "widget_ids": {
        "computed": true,
        "description": "The list of widget identifiers included in the scheduled report. If not specified, all widgets in the dashboard are included.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Creates and manages a scheduled report for an AWS Billing and Cost Management Dashboards dashboard. A scheduled report renders a snapshot of the dashboard on a cron schedule, using an IAM role that the service assumes to read the underlying cost and usage data.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBcmScheduledReportSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBcmScheduledReport), &result)
	return &result
}
