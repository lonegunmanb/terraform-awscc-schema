package resource

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
        "optional": true,
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
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "index_id": {
        "description": "Index ID",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "language_code": {
        "computed": true,
        "description": "The code for a language.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "FAQ name",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_arn": {
        "description": "FAQ role ARN",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "s3_path": {
        "description": "FAQ S3 path",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "bucket": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "Tags for labeling the FAQ",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "A string used to identify this tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "A string containing the value for the tag",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "A Kendra FAQ resource",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccKendraFaqSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccKendraFaq), &result)
	return &result
}
