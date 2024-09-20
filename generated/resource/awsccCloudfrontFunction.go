package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontFunction = `{
  "block": {
    "attributes": {
      "auto_publish": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_config": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_value_store_associations": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_value_store_arn": {
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
            "runtime": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "function_metadata": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "function_arn": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::Function",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFunction), &result)
	return &result
}
