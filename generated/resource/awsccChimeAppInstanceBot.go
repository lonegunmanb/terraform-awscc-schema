package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccChimeAppInstanceBot = `{
  "block": {
    "attributes": {
      "app_instance_arn": {
        "description": "The ARN of the AppInstance.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_instance_bot_arn": {
        "computed": true,
        "description": "The ARN of the AppInstanceBot.",
        "description_kind": "plain",
        "type": "string"
      },
      "configuration": {
        "description": "A structure that contains configuration data.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "lex": {
              "description": "The configuration for an Amazon Lex V2 bot.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "invoked_by": {
                    "computed": true,
                    "description": "Specifies the type of message that triggers a bot.",
                    "description_kind": "plain",
                    "nested_type": {
                      "attributes": {
                        "standard_messages": {
                          "computed": true,
                          "description": "Sets standard messages as the bot trigger.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "targeted_messages": {
                          "computed": true,
                          "description": "Sets targeted messages as the bot trigger.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "nesting_mode": "single"
                    },
                    "optional": true
                  },
                  "lex_bot_alias_arn": {
                    "description": "The ARN of the Amazon Lex V2 bot's alias.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "locale_id": {
                    "description": "Identifies the Amazon Lex V2 bot's language and locale.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "responds_to": {
                    "computed": true,
                    "description": "Determines whether the Amazon Lex V2 bot responds to all standard messages. Control messages are not supported.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "welcome_intent": {
                    "computed": true,
                    "description": "The name of the welcome intent configured in the Amazon Lex V2 bot.",
                    "description_kind": "plain",
                    "optional": true,
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
      "created_timestamp": {
        "computed": true,
        "description": "The time at which the AppInstanceBot was created. In epoch milliseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_timestamp": {
        "computed": true,
        "description": "The time at which the AppInstanceBot was last updated. In epoch milliseconds.",
        "description_kind": "plain",
        "type": "number"
      },
      "metadata": {
        "computed": true,
        "description": "The metadata of the AppInstanceBot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the AppInstanceBot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The tags assigned to the AppInstanceBot.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key in a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value in a tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::Chime::AppInstanceBot",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccChimeAppInstanceBotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccChimeAppInstanceBot), &result)
	return &result
}
