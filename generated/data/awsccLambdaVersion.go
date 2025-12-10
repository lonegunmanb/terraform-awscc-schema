package data

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
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A description for the version to override the description in the function configuration. Updates are not supported for this property.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_arn": {
        "computed": true,
        "description": "The ARN of the version.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_name": {
        "computed": true,
        "description": "The name of the Lambda function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_scaling_config": {
        "computed": true,
        "description": "The scaling configuration to apply to the function, including minimum and maximum execution environment limits.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_execution_environments": {
              "computed": true,
              "description": "The maximum number of execution environments that can be provisioned for the function.",
              "description_kind": "plain",
              "type": "number"
            },
            "min_execution_environments": {
              "computed": true,
              "description": "The minimum number of execution environments to maintain for the function.",
              "description_kind": "plain",
              "type": "number"
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
      "provisioned_concurrency_config": {
        "computed": true,
        "description": "Specifies a provisioned concurrency configuration for a function's version. Updates are not supported for this property.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "provisioned_concurrent_executions": {
              "computed": true,
              "description": "The amount of provisioned concurrency to allocate for the version.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "update_runtime_on": {
              "computed": true,
              "description": "The runtime update mode.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "version": {
        "computed": true,
        "description": "The version number.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Lambda::Version",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLambdaVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaVersion), &result)
	return &result
}
