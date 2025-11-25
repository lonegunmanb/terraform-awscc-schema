package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudfrontConnectionFunction = `{
  "block": {
    "attributes": {
      "auto_publish": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "connection_function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_function_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_function_config": {
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
      "connection_function_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "e_tag": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CloudFront::ConnectionFunction",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudfrontConnectionFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontConnectionFunction), &result)
	return &result
}
