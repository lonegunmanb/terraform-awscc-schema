package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMacieAllowList = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "AllowList ARN.",
        "description_kind": "plain",
        "type": "string"
      },
      "criteria": {
        "computed": true,
        "description": "AllowList criteria.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "regex": {
              "computed": true,
              "description": "The S3 object key for the AllowList.",
              "description_kind": "plain",
              "type": "string"
            },
            "s3_words_list": {
              "computed": true,
              "description": "The S3 location for the AllowList.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "bucket_name": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "object_key": {
                    "computed": true,
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
        "description": "Description of AllowList.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Name of AllowList.",
        "description_kind": "plain",
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
              "computed": true,
              "description": "The tag's key.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The tag's value.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::Macie::AllowList",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMacieAllowListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMacieAllowList), &result)
	return &result
}
