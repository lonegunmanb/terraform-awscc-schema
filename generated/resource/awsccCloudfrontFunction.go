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
        "description": "A flag that determines whether to automatically publish the function to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage when it?s created. To automatically publish to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage, set this property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
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
        "description": "The function code. For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide*.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_config": {
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
        "description": "A name to identify the function.",
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
    "description": "Creates a CF function.\n To create a function, you provide the function code and some configuration information about the function. The response contains an Amazon Resource Name (ARN) that uniquely identifies the function, and the function?s stage.\n By default, when you create a function, it?s in the ` + "`" + `` + "`" + `DEVELOPMENT` + "`" + `` + "`" + ` stage. In this stage, you can [test the function](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/test-function.html) in the CF console (or with ` + "`" + `` + "`" + `TestFunction` + "`" + `` + "`" + ` in the CF API).\n When you?re ready to use your function with a CF distribution, publish the function to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage. You can do this in the CF console, with ` + "`" + `` + "`" + `PublishFunction` + "`" + `` + "`" + ` in the CF API, or by updating the ` + "`" + `` + "`" + `AWS::CloudFront::Function` + "`" + `` + "`" + ` resource with the ` + "`" + `` + "`" + `AutoPublish` + "`" + `` + "`" + ` property set to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `. When the function is published to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage, you can attach it to a distribution?s cache behavior, using the function?s ARN.\n To automatically publish the function to the ` + "`" + `` + "`" + `LIVE` + "`" + `` + "`" + ` stage when it?s created, set the ` + "`" + `` + "`" + `AutoPublish` + "`" + `` + "`" + ` property to ` + "`" + `` + "`" + `true` + "`" + `` + "`" + `.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudfrontFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudfrontFunction), &result)
	return &result
}
