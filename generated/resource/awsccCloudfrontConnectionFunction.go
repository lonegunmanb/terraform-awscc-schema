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
        "description": "A flag that determines whether to automatically publish the function to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage when it?s created. To automatically publish to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage, set this property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
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
        "description": "The code for the connection function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "connection_function_config": {
        "description": "Contains configuration information about a CloudFront function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "description": "A comment to describe the function.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "key_value_store_associations": {
              "computed": true,
              "description": "The configuration for the key value store associations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key_value_store_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the key value store association.",
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
              "description": "The function's runtime environment version.",
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
        "description": "The connection function name.",
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
        "description": "A complex type that contains zero or more ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` elements.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A string that contains ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` key.\n The string length should be between 1 and 128 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "A string that contains an optional ` + "`" + `` + "`" + `Tag` + "`" + `` + "`" + ` value.\n The string length should be between 0 and 256 characters. Valid characters include ` + "`" + `` + "`" + `a-z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `A-Z` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `0-9` + "`" + `` + "`" + `, space, and the special characters ` + "`" + `` + "`" + `_ - . : / = + @` + "`" + `` + "`" + `.",
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
    "description": "A connection function.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontConnectionFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontConnectionFunction), &result)
	return &result
}
