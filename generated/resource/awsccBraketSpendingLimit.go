package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccBraketSpendingLimit = `{
  "block": {
    "attributes": {
      "created_at": {
        "computed": true,
        "description": "The date and time when the spending limit was created, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      },
      "device_arn": {
        "description": "The Amazon Resource Name (ARN) of the quantum device to apply the spending limit to.",
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
      "queued_spend": {
        "computed": true,
        "description": "The amount currently queued for spending on the device, in USD.",
        "description_kind": "plain",
        "type": "string"
      },
      "spending_limit": {
        "description": "The maximum amount that can be spent on the specified device, in USD.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "spending_limit_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that uniquely identifies the spending limit.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags to apply to the spending limit.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "time_period": {
        "computed": true,
        "description": "Defines a time range for spending limits, specifying when the limit is active.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "end_at": {
              "computed": true,
              "description": "The end date and time for the spending limit period, in ISO 8601 format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "start_at": {
              "computed": true,
              "description": "The start date and time for the spending limit period, in ISO 8601 format.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "total_spend": {
        "computed": true,
        "description": "The total amount spent on the device so far during the current time period, in USD.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The date and time when the spending limit was last modified, in ISO 8601 format.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a spending limit for a specified quantum device. Spending limits help you control costs by setting maximum amounts that can be spent on quantum computing tasks within a specified time period.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccBraketSpendingLimitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccBraketSpendingLimit), &result)
	return &result
}
