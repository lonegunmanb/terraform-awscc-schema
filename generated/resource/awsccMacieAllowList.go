package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieAllowList = `{
  "block": {
    "attributes": {
      "allow_list_id": {
        "computed": true,
        "description": "AllowList ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "arn": {
        "computed": true,
        "description": "AllowList ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "criteria": {
        "description": "AllowList criteria.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "regex": {
              "computed": true,
              "description": "The S3 object key for the AllowList.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "s3_words_list": {
              "computed": true,
              "description": "The S3 location for the AllowList.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "object_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "description": {
        "computed": true,
        "description": "Description of AllowList.",
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
      "name": {
        "description": "Name of AllowList.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "AllowList status.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A collection of tags associated with a resource",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The tag's key.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The tag's value.",
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
    "description": "Macie AllowList resource schema",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMacieAllowListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieAllowList), &result)
	return &result
}
