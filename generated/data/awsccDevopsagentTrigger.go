package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentTrigger = `{
  "block": {
    "attributes": {
      "action": {
        "computed": true,
        "description": "The action to perform when the trigger fires. A JSON object containing actionType and task.",
        "description_kind": "plain",
        "type": "string"
      },
      "agent_space_id": {
        "computed": true,
        "description": "The unique identifier of the parent Agent Space.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the trigger. Nested under the parent Agent Space: arn:\u003cpartition\u003e:aidevops:\u003cregion\u003e:\u003caccount-id\u003e:agentspace/\u003cagentspace-id\u003e/trigger/\u003ctrigger-id\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "condition": {
        "computed": true,
        "description": "The condition that causes the trigger to fire.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule": {
              "computed": true,
              "description": "Schedule configuration for a time-based trigger.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "computed": true,
                    "description": "A cron or rate expression that defines when the trigger fires.",
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
      "created_at": {
        "computed": true,
        "description": "The timestamp when the trigger was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the trigger. Active triggers fire on schedule; Inactive triggers are paused.",
        "description_kind": "plain",
        "type": "string"
      },
      "trigger_id": {
        "computed": true,
        "description": "The unique identifier of the trigger (assigned by the service on Create).",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description": "The type of trigger.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the trigger was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DevOpsAgent::Trigger",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsagentTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentTrigger), &result)
	return &result
}
