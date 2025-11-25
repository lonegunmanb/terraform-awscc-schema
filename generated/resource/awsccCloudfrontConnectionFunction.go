package resource

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
        "optional": true,
        "type": "bool"
      },
      "connection_function_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_function_code": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_function_config": {
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
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "last_modified_time": {
        "computed": true,
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
      }
    },
    "description": "Resource Type definition for AWS::CloudFront::ConnectionFunction",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontConnectionFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontConnectionFunction), &result)
	return &result
}
