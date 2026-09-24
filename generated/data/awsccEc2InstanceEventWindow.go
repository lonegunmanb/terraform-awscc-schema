package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccEc2InstanceEventWindow = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the event window.",
        "description_kind": "plain",
        "type": "string"
      },
      "cron_expression": {
        "computed": true,
        "description": "The cron expression defined for the event window. Exactly one of TimeRanges or CronExpression must be specified.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_event_window_id": {
        "computed": true,
        "description": "The ID of the event window.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the event window.",
        "description_kind": "plain",
        "type": "string"
      },
      "state": {
        "computed": true,
        "description": "The current state of the event window. Observed values include creating, active, modifying, deleting and deleted.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags applied to the event window.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "time_ranges": {
        "computed": true,
        "description": "The time ranges of the event window. Exactly one of TimeRanges or CronExpression must be specified.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_hour": {
              "computed": true,
              "description": "The hour when the time range ends.",
              "description_kind": "plain",
              "type": "number"
            },
            "end_week_day": {
              "computed": true,
              "description": "The day on which the time range ends.",
              "description_kind": "plain",
              "type": "string"
            },
            "start_hour": {
              "computed": true,
              "description": "The hour when the time range begins.",
              "description_kind": "plain",
              "type": "number"
            },
            "start_week_day": {
              "computed": true,
              "description": "The day on which the time range begins.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::EC2::InstanceEventWindow",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccEc2InstanceEventWindowSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccEc2InstanceEventWindow), &result)
	return &result
}
