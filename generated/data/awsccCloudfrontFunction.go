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
        "description": "A flag that determines whether to automatically publish the function to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage when it?s created. To automatically publish to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage, set this property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
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
        "description": "The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_config": {
        "computed": true,
        "description": "Contains configuration information about a CloudFront function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "comment": {
              "computed": true,
              "description": "A comment to describe the function.",
              "description_kind": "plain",
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
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "runtime": {
              "computed": true,
              "description": "The function's runtime environment version.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "function_metadata": {
        "computed": true,
        "description": "Contains metadata about a CloudFront function.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "function_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the function. The ARN uniquely identifies the function.",
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
        "description": "A name to identify the function.",
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
