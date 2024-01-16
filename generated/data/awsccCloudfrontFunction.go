package data

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
        "type": "bool"
      },
      "function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "function_config": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description_kind": "plain",
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
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "runtime": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
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
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudFront::Function",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFunction), &result)
	return &result
}
