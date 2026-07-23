package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language_code": {
        "computed": true,
        "description": "The language code that represents the language of the entries in your vocabulary filter.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value of the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "vocabulary_filter_file_uri": {
        "computed": true,
        "description": "The Amazon S3 location of the text file that contains your custom vocabulary filter terms.",
        "description_kind": "plain",
        "type": "string"
      },
      "vocabulary_filter_name": {
        "computed": true,
        "description": "A unique name, chosen by you, for your custom vocabulary filter.",
        "description_kind": "plain",
        "type": "string"
      },
      "words": {
        "computed": true,
        "description": "Use this parameter if you want to create your custom vocabulary filter by including all desired terms, as comma-separated values, within your request.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::Transcribe::VocabularyFilter",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccTranscribeVocabularyFilterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranscribeVocabularyFilter), &result)
	return &result
}
