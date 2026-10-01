package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranscribeMedicalVocabulary = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the medical vocabulary.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "language_code": {
        "description": "The language code of the vocabulary entries.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the medical vocabulary.",
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
        "description": "The Amazon S3 location of the text file that contains the medical vocabulary.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vocabulary_name": {
        "description": "The name of the medical vocabulary.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::Transcribe::MedicalVocabulary",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTranscribeMedicalVocabularySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranscribeMedicalVocabulary), &result)
	return &result
}
