package resource

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
        "description": "Members of the rotation",
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "Name of the Rotation",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "recurrence": {
        "description": "Information about when an on-call rotation is in effect and how long the rotation period lasts.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "daily_settings": {
              "computed": true,
              "description": "Information about on-call rotations that recur daily.",
              "description_kind": "plain",
              "optional": true,
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
                    "description": "The day of the month when monthly recurring on-call rotations begin.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "hand_off_time": {
                    "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "number_of_on_calls": {
              "computed": true,
              "description": "Number of Oncalls per shift.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "recurrence_multiplier": {
              "computed": true,
              "description": "The number of days, weeks, or months a single rotation lasts.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "shift_coverages": {
              "computed": true,
              "description": "Information about the days of the week included in on-call rotation coverage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "coverage_times": {
                    "description": "Information about when an on-call shift begins and ends.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "end_time": {
                          "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "start_time": {
                          "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "list"
                    },
                    "required": true
                  },
                  "day_of_week": {
                    "description": "The day of the week when weekly recurring on-call shift rotations begin. ",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "weekly_settings": {
              "computed": true,
              "description": "Information about on-call rotations that recur weekly.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "day_of_week": {
                    "description": "The day of the week when weekly recurring on-call shift rotations begin. ",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "hand_off_time": {
                    "description": "Details about when an on-call rotation shift begins or ends. Time of the day in format HH:MM",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "start_time": {
        "description": "Start time of the first shift of Oncall Schedule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "time_zone_id": {
        "description": "TimeZone Identifier for the Oncall Schedule",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::SSMContacts::Rotation.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmcontactsRotationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmcontactsRotation), &result)
	return &result
}
