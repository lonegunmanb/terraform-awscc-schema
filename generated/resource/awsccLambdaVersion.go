package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaVersion = `{
  "block": {
    "attributes": {
      "code_sha_256": {
        "computed": true,
        "description": "Only publish a version if the hash value matches the value that's specified. Use this option to avoid publishing a version if the function code has changed since you last updated it. Updates are not supported for this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the version to override the description in the function configuration. Updates are not supported for this property.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_arn": {
        "computed": true,
        "description": "The ARN of the version.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "description": "The name of the Lambda function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "provisioned_concurrency_config": {
        "computed": true,
        "description": "Specifies a provisioned concurrency configuration for a function's version. Updates are not supported for this property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provisioned_concurrent_executions": {
              "description": "The amount of provisioned concurrency to allocate for the version.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "runtime_policy": {
        "computed": true,
        "description": "Specifies the runtime management configuration of a function. Displays runtimeVersionArn only for Manual.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "runtime_version_arn": {
              "computed": true,
              "description": "The ARN of the runtime the function is configured to use. If the runtime update mode is manual, the ARN is returned, otherwise null is returned.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update_runtime_on": {
              "description": "The runtime update mode.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "version": {
        "computed": true,
        "description": "The version number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::Version",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaVersion), &result)
	return &result
}
