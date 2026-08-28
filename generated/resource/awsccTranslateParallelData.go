package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccTranslateParallelData = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the parallel data resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time at which the parallel data resource was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A custom description for the parallel data resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encryption_key": {
        "computed": true,
        "description": "The encryption key used to encrypt this object.",
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
      "failed_record_count": {
        "computed": true,
        "description": "The number of records unsuccessfully imported.",
        "description_kind": "plain",
        "type": "number"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "imported_data_size": {
        "computed": true,
        "description": "The number of UTF-8 characters imported from the parallel data input file.",
        "description_kind": "plain",
        "type": "number"
      },
      "imported_record_count": {
        "computed": true,
        "description": "The number of records successfully imported.",
        "description_kind": "plain",
        "type": "number"
      },
      "last_updated_at": {
        "computed": true,
        "description": "The time at which the parallel data resource was last updated.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "A custom name for the parallel data resource. Must be unique in the account and region.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parallel_data_config": {
        "description": "Specifies the format and S3 location of the parallel data input file.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "format": {
              "description": "The format of the parallel data input file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "s3_uri": {
              "description": "The URI of the Amazon S3 folder that contains the parallel data input file.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "skipped_record_count": {
        "computed": true,
        "description": "The number of items skipped during import.",
        "description_kind": "plain",
        "type": "number"
      },
      "source_language_code": {
        "computed": true,
        "description": "The source language of the translations in the parallel data file.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The status of the parallel data resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "Tags associated with the parallel data resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
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
        "description": "The language codes for the target languages available in the parallel data file.",
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      }
    },
    "description": "A parallel data resource in Amazon Translate used to customize machine translation output.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccTranslateParallelDataSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccTranslateParallelData), &result)
	return &result
}
