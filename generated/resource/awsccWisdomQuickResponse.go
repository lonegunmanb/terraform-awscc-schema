package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccWisdomQuickResponse = `{
  "block": {
    "attributes": {
      "channels": {
        "computed": true,
        "description": "The Amazon Connect contact channels this quick response applies to.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "content": {
        "description": "The container of quick response content.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "content": {
              "computed": true,
              "description": "The content of the quick response.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "content_type": {
        "computed": true,
        "description": "The media type of the quick response content.\n- Use application/x.quickresponse;format=plain for quick response written in plain text.\n- Use application/x.quickresponse;format=markdown for quick response written in richtext.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "contents": {
        "computed": true,
        "description": "The content of the quick response stored in different media types.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "markdown": {
              "computed": true,
              "description": "The container of quick response content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content": {
                    "computed": true,
                    "description": "The content of the quick response.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "plain_text": {
              "computed": true,
              "description": "The container of quick response content.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "content": {
                    "computed": true,
                    "description": "The content of the quick response.",
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
      "description": {
        "computed": true,
        "description": "The description of the quick response.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "grouping_configuration": {
        "computed": true,
        "description": "The configuration information of the user groups that the quick response is accessible to.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "criteria": {
              "computed": true,
              "description": "The criteria used for grouping Amazon Q in Connect users.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "values": {
              "computed": true,
              "description": "The list of values that define different groups of Amazon Q in Connect users.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "is_active": {
        "computed": true,
        "description": "Whether the quick response is active.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "knowledge_base_arn": {
        "description": "The Amazon Resource Name (ARN) of the knowledge base.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language": {
        "computed": true,
        "description": "The language code value for the language in which the quick response is written. The supported language codes include de_DE, en_US, es_ES, fr_FR, id_ID, it_IT, ja_JP, ko_KR, pt_BR, zh_CN, zh_TW",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the quick response.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "quick_response_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the quick response.",
        "description_kind": "plain",
        "type": "string"
      },
      "quick_response_id": {
        "computed": true,
        "description": "The identifier of the quick response.",
        "description_kind": "plain",
        "type": "string"
      },
      "shortcut_key": {
        "computed": true,
        "description": "The shortcut key of the quick response. The value should be unique across the knowledge base.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the quick response data.",
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
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::Wisdom::QuickResponse Resource Type.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccWisdomQuickResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccWisdomQuickResponse), &result)
	return &result
}
