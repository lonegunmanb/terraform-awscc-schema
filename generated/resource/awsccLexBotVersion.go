package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLexBotVersion = `{
  "block": {
    "attributes": {
      "bot_id": {
        "description": "Unique ID of resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "bot_version": {
        "computed": true,
        "description": "The version of a bot.",
        "description_kind": "plain",
        "type": "string"
      },
      "bot_version_locale_specification": {
        "description": "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bot_version_locale_details": {
              "description": "The version of a bot used for a bot locale.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "source_bot_version": {
                    "description": "The version of a bot.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "required": true
            },
            "locale_id": {
              "description": "The identifier of the language and locale that the bot will be used in.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "A description of the version. Use the description to help identify the version in lists.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "A version is a numbered snapshot of your work that you can publish for use in different parts of your workflow, such as development, beta deployment, and production.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLexBotVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLexBotVersion), &result)
	return &result
}
