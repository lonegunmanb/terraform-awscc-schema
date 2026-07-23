package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranscribeVocabularyFilter = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the vocabulary filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "language_code": {
        "description": "The language code that represents the language of the entries in your vocabulary filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the vocabulary filter.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "vocabulary_filter_file_uri": {
        "computed": true,
        "description": "The Amazon S3 location of the text file that contains your custom vocabulary filter terms.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vocabulary_filter_name": {
        "description": "A unique name, chosen by you, for your custom vocabulary filter.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "words": {
        "computed": true,
        "description": "Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Creates a custom vocabulary filter that you can use to mask, delete, or flag specific words from your transcript.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTranscribeVocabularyFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranscribeVocabularyFilter), &result)
	return &result
}
