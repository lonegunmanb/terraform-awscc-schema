package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccQuicksightRefreshSchedule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "\u003cp\u003eThe Amazon Resource Name (ARN) of the data source.\u003c/p\u003e",
        "description_kind": "plain",
        "type": "string"
      },
      "aws_account_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_set_id": {
        "computed": true,
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
      "schedule": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "refresh_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "schedule_frequency": {
              "computed": true,
              "description": "\u003cp\u003eInformation about the schedule frequency.\u003c/p\u003e",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "interval": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "refresh_on_day": {
                    "computed": true,
                    "description": "\u003cp\u003eThe day scheduled for refresh.\u003c/p\u003e",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "day_of_month": {
                          "computed": true,
                          "description": "\u003cp\u003eThe Day Of Month for scheduled refresh.\u003c/p\u003e",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "day_of_week": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "time_of_the_day": {
                    "computed": true,
                    "description": "\u003cp\u003eThe time of the day for scheduled refresh.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "time_zone": {
                    "computed": true,
                    "description": "\u003cp\u003eThe timezone for scheduled refresh.\u003c/p\u003e",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "schedule_id": {
              "computed": true,
              "description": "\u003cp\u003eAn unique identifier for the refresh schedule.\u003c/p\u003e",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_after_date_time": {
              "computed": true,
              "description": "\u003cp\u003eThe date time after which refresh is to be scheduled\u003c/p\u003e",
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
    "description": "Definition of the AWS::QuickSight::RefreshSchedule Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccQuicksightRefreshScheduleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccQuicksightRefreshSchedule), &result)
	return &result
}
