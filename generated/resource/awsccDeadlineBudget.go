package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineBudget = `{
  "block": {
    "attributes": {
      "actions": {
        "description": "The budget actions to specify what happens when the budget runs out.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description for the budget action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "threshold_percentage": {
              "description": "The percentage threshold for the budget action.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "type": {
              "description": "The type of budget action.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "approximate_dollar_limit": {
        "description": "The dollar limit based on consumed usage.",
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "arn": {
        "computed": true,
        "description": "The ARN of the budget.",
        "description_kind": "plain",
        "type": "string"
      },
      "budget_id": {
        "computed": true,
        "description": "The budget ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the budget.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "The display name of the budget.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "farm_id": {
        "description": "The farm ID.",
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
      "schedule": {
        "description": "The start and end time of the budget.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fixed": {
              "description": "The details of a fixed budget schedule.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time": {
                    "description": "When the budget ends.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "start_time": {
                    "description": "When the budget starts.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "status": {
        "computed": true,
        "description": "The status of the budget.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
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
          "nesting_mode": "list"
        },
        "optional": true
      },
      "usage_tracking_resource": {
        "description": "The usage details of the allotted budget.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "queue_id": {
              "description": "The queue ID.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      }
    },
    "description": "Creates a budget to set spending thresholds for your rendering activity.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDeadlineBudgetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineBudget), &result)
	return &result
}
