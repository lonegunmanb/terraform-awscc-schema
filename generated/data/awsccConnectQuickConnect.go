package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectQuickConnect = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the quick connect.",
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
        "description": "The name of the quick connect.",
        "description_kind": "plain",
        "type": "string"
      },
      "quick_connect_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the quick connect.",
        "description_kind": "plain",
        "type": "string"
      },
      "quick_connect_config": {
        "computed": true,
        "description": "Configuration settings for the quick connect.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "phone_config": {
              "computed": true,
              "description": "The phone configuration. This is required only if QuickConnectType is PHONE_NUMBER.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "phone_number": {
                    "computed": true,
                    "description": "The phone number in E.164 format.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "queue_config": {
              "computed": true,
              "description": "The queue configuration. This is required only if QuickConnectType is QUEUE.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "contact_flow_arn": {
                    "computed": true,
                    "description": "The identifier of the contact flow.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "queue_arn": {
                    "computed": true,
                    "description": "The identifier for the queue.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "quick_connect_type": {
              "computed": true,
              "description": "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
              "description_kind": "plain",
              "type": "string"
            },
            "user_config": {
              "computed": true,
              "description": "The user configuration. This is required only if QuickConnectType is USER.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "contact_flow_arn": {
                    "computed": true,
                    "description": "The identifier of the contact flow.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "user_arn": {
                    "computed": true,
                    "description": "The identifier of the user.",
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
      "quick_connect_type": {
        "computed": true,
        "description": "The type of quick connect. In the Amazon Connect console, when you create a quick connect, you are prompted to assign one of the following types: Agent (USER), External (PHONE_NUMBER), or Queue (QUEUE).",
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
      }
    },
    "description": "Data Source schema for AWS::Connect::QuickConnect",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectQuickConnectSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectQuickConnect), &result)
	return &result
}
