package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranslateTerminology = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the custom terminology.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the terminology was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the custom terminology.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "directionality": {
        "computed": true,
        "description": "The directionality of the terminology.",
        "description_kind": "plain",
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "The encryption key for the custom terminology.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "id": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the encryption key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The type of encryption key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "format": {
        "computed": true,
        "description": "The format of the terminology data.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The time the terminology was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "merge_strategy": {
        "computed": true,
        "description": "The merge strategy for the custom terminology. Currently only OVERWRITE is supported.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the custom terminology.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "size_bytes": {
        "computed": true,
        "description": "The size of the terminology file in bytes.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_language_code": {
        "computed": true,
        "description": "The source language code for the custom terminology.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the terminology.",
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
      "target_language_codes": {
        "computed": true,
        "description": "The target language codes for the custom terminology.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "term_count": {
        "computed": true,
        "description": "The number of terms in the custom terminology.",
        "description_kind": "plain",
        "type": "number"
      },
      "terminology_data": {
        "computed": true,
        "description": "The terminology data for the custom terminology being imported.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "directionality": {
              "computed": true,
              "description": "The directionality of the terminology resource.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "file": {
              "computed": true,
              "description": "The file containing the custom terminology data, base64-encoded.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format": {
              "computed": true,
              "description": "The data format of the custom terminology.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "A custom terminology resource for Amazon Translate that enables customized translations.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTranslateTerminologySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranslateTerminology), &result)
	return &result
}
