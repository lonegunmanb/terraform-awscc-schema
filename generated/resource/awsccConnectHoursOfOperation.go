package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectHoursOfOperation = `{
  "block": {
    "attributes": {
      "config": {
        "description": "Configuration information for the hours of operation: day, start time, and end time.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "day": {
              "description": "The day that the hours of operation applies to.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "end_time": {
              "description": "The end time that your contact center closes.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "description": "The hours.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "start_time": {
              "description": "The start time that your contact center opens.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "hours": {
                    "description": "The hours.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "minutes": {
                    "description": "The minutes.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "set"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "The description of the hours of operation.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "hours_of_operation_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the hours of operation.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "instance_arn": {
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the hours of operation.",
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "time_zone": {
        "description": "The time zone of the hours of operation.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Connect::HoursOfOperation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccConnectHoursOfOperationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectHoursOfOperation), &result)
	return &result
}
