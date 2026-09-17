package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranscribeVocabulary = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom vocabulary.",
        "description_kind": "plain",
        "type": "string"
      },
      "data_access_role_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input file.",
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
        "description": "The language code that represents the language of the entries in your custom vocabulary. Each custom vocabulary must contain terms in only one language.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description": "The date and time the specified custom vocabulary was last modified.",
        "description_kind": "plain",
        "type": "string"
      },
      "phrases": {
        "computed": true,
        "description": "Use this parameter if you want to create your custom vocabulary by including all desired terms, as comma-separated values, within your request. You cannot use this parameter together with VocabularyFileUri.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "Adds one or more custom tags, each in the form of a key:value pair, to the custom vocabulary.",
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
      "vocabulary_file_uri": {
        "computed": true,
        "description": "The Amazon S3 location of the text file that contains your custom vocabulary. You cannot use this parameter together with Phrases.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vocabulary_name": {
        "description": "A unique name, chosen by you, for your custom vocabulary. This name is case sensitive, cannot contain spaces, and must be unique within an AWS account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vocabulary_state": {
        "computed": true,
        "description": "The processing state of your custom vocabulary. If the state is READY, you can use the custom vocabulary in a StartTranscriptionJob request.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Creates a custom vocabulary that you can use to improve the transcription accuracy of domain-specific words and phrases.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTranscribeVocabularySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranscribeVocabulary), &result)
	return &result
}
