package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChatbotCustomAction = `{
  "block": {
    "attributes": {
      "action_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "alias_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "attachments": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "button_text": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "criteria": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "operator": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "variable_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "notification_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "variables": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        }
      },
      "custom_action_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "definition": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "command_text": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
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
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Chatbot::CustomAction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccChatbotCustomActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChatbotCustomAction), &result)
	return &result
}
