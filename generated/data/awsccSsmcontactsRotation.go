package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmcontactsRotation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the rotation.",
        "description_kind": "plain",
        "type": "string"
      },
      "contact_ids": {
        "computed": true,
        "description": "Members of the rotation",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of the Rotation",
        "description_kind": "plain",
        "type": "string"
      },
      "recurrence": {
        "computed": true,
        "description": "Information about when an on-call rotation is in effect and how long the rotation period lasts.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "daily_settings": {
              "computed": true,
              "description": "Information about on-call rotations that recur daily.",
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "monthly_settings": {
              "computed": true,
              "description": "Information about on-call rotations that recur monthly.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day_of_month": {
                    "computed": true,
                    "description": "The day of the month when monthly recurring on-call rotations begin.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "hand_off_time": {
                    "computed": true,
                    "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "number_of_on_calls": {
              "computed": true,
              "description": "Number of Oncalls per shift.",
              "description_kind": "plain",
              "type": "number"
            },
            "recurrence_multiplier": {
              "computed": true,
              "description": "The number of days, weeks, or months a single rotation lasts.",
              "description_kind": "plain",
              "type": "number"
            },
            "shift_coverages": {
              "computed": true,
              "description": "Information about the days of the week included in on-call rotation coverage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "coverage_times": {
                    "computed": true,
                    "description": "Information about when an on-call shift begins and ends.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "computed": true,
                          "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                          "description_kind": "plain",
                          "type": "string"
                        },
                        "start_time": {
                          "computed": true,
                          "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                          "description_kind": "plain",
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    }
                  },
                  "day_of_week": {
                    "computed": true,
                    "description": "The day of the week when weekly recurring on-call shift rotations begin. ",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "weekly_settings": {
              "computed": true,
              "description": "Information about on-call rotations that recur weekly.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day_of_week": {
                    "computed": true,
                    "description": "The day of the week when weekly recurring on-call shift rotations begin. ",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "hand_off_time": {
                    "computed": true,
                    "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "start_time": {
        "computed": true,
        "description": "Start time of the first shift of Oncall Schedule",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "time_zone_id": {
        "computed": true,
        "description": "TimeZone Identifier for the Oncall Schedule",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSMContacts::Rotation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmcontactsRotationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmcontactsRotation), &result)
	return &result
}
