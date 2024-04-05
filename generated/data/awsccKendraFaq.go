package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccKendraFaq = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "FAQ description",
        "description_kind": "plain",
        "type": "string"
      },
      "faq_id": {
        "computed": true,
        "description": "Unique ID of the FAQ",
        "description_kind": "plain",
        "type": "string"
      },
      "file_format": {
        "computed": true,
        "description": "FAQ file format",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "index_id": {
        "computed": true,
        "description": "Index ID",
        "description_kind": "plain",
        "type": "string"
      },
      "language_code": {
        "computed": true,
        "description": "The code for a language.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "FAQ name",
        "description_kind": "plain",
        "type": "string"
      },
      "role_arn": {
        "computed": true,
        "description": "FAQ role ARN",
        "description_kind": "plain",
        "type": "string"
      },
      "s3_path": {
        "computed": true,
        "description": "FAQ S3 path",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "Tags for labeling the FAQ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Kendra::Faq",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccKendraFaqSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraFaq), &result)
	return &result
}
