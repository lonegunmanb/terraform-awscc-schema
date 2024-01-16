package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectHoursOfOperation = `{
  "block": {
    "attributes": {
      "config": {
        "computed": true,
        "description": "Configuration information for the hours of operation: day, start time, and end time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "day": {
              "computed": true,
              "description": "The day that the hours of operation applies to.",
              "description_kind": "plain",
              "type": "string"
            },
            "end_time": {
              "computed": true,
              "description": "The end time that your contact center closes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "computed": true,
                    "description": "The hours.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "minutes": {
                    "computed": true,
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "start_time": {
              "computed": true,
              "description": "The start time that your contact center opens.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "computed": true,
                    "description": "The hours.",
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "minutes": {
                    "computed": true,
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "set"
        }
      },
      "description": {
        "computed": true,
        "description": "The description of the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "hours_of_operation_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "time_zone": {
        "computed": true,
        "description": "The time zone of the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Connect::HoursOfOperation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectHoursOfOperationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectHoursOfOperation), &result)
	return &result
}
