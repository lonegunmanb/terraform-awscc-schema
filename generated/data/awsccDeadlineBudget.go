package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDeadlineBudget = `{
  "block": {
    "attributes": {
      "actions": {
        "computed": true,
        "description": "The budget actions to specify what happens when the budget runs out.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "description": {
              "computed": true,
              "description": "A description for the budget action.",
              "description_kind": "plain",
              "type": "string"
            },
            "threshold_percentage": {
              "computed": true,
              "description": "The percentage threshold for the budget action.",
              "description_kind": "plain",
              "type": "number"
            },
            "type": {
              "computed": true,
              "description": "The type of budget action.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "approximate_dollar_limit": {
        "computed": true,
        "description": "The dollar limit based on consumed usage.",
        "description_kind": "plain",
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
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The display name of the budget.",
        "description_kind": "plain",
        "type": "string"
      },
      "farm_id": {
        "computed": true,
        "description": "The farm ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schedule": {
        "computed": true,
        "description": "The start and end time of the budget.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fixed": {
              "computed": true,
              "description": "The details of a fixed budget schedule.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "end_time": {
                    "computed": true,
                    "description": "When the budget ends.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "start_time": {
                    "computed": true,
                    "description": "When the budget starts.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
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
      "usage_tracking_resource": {
        "computed": true,
        "description": "The usage details of the allotted budget.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "queue_id": {
              "computed": true,
              "description": "The queue ID.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::Deadline::Budget",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDeadlineBudgetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDeadlineBudget), &result)
	return &result
}
