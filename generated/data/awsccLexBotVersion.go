package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLexBotVersion = `{
  "block": {
    "attributes": {
      "bot_id": {
        "computed": true,
        "description": "Unique ID of resource",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_version": {
        "computed": true,
        "description": "The version of a bot.",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_version_locale_specification": {
        "computed": true,
        "description": "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bot_version_locale_details": {
              "computed": true,
              "description": "The version of a bot used for a bot locale.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_bot_version": {
                    "computed": true,
                    "description": "The version of a bot.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "locale_id": {
              "computed": true,
              "description": "The identifier of the language and locale that the bot will be used in.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the version. Use the description to help identify the version in lists.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lex::BotVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLexBotVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLexBotVersion), &result)
	return &result
}
