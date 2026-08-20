package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsagentTrigger = `{
  "block": {
    "attributes": {
      "action": {
        "description": "The action to perform when the trigger fires. A JSON object containing actionType and task.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "agent_space_id": {
        "description": "The unique identifier of the parent Agent Space.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the trigger. Nested under the parent Agent Space: arn:\u003cpartition\u003e:aidevops:\u003cregion\u003e:\u003caccount-id\u003e:agentspace/\u003cagentspace-id\u003e/trigger/\u003ctrigger-id\u003e.",
        "description_kind": "plain",
        "type": "string"
      },
      "condition": {
        "description": "The condition that causes the trigger to fire.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "schedule": {
              "description": "Schedule configuration for a time-based trigger.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "expression": {
                    "description": "A cron or rate expression that defines when the trigger fires.",
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
      "created_at": {
        "computed": true,
        "description": "The timestamp when the trigger was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the trigger. Active triggers fire on schedule; Inactive triggers are paused.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trigger_id": {
        "computed": true,
        "description": "The unique identifier of the trigger (assigned by the service on Create).",
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "description": "The type of trigger.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "updated_at": {
        "computed": true,
        "description": "The timestamp when the trigger was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::DevOpsAgent::Trigger. A trigger defines an automated action that fires on a schedule within an Agent Space.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsagentTriggerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsagentTrigger), &result)
	return &result
}
