package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaEventInvokeConfig = `{
  "block": {
    "attributes": {
      "destination_config": {
        "computed": true,
        "description": "A destination for events after they have been sent to a function for processing.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "on_failure": {
              "computed": true,
              "description": "The destination configuration for failed invocations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the destination resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "on_success": {
              "computed": true,
              "description": "The destination configuration for successful invocations.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the destination resource.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
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
      "maximum_event_age_in_seconds": {
        "computed": true,
        "description": "The maximum age of a request that Lambda sends to a function for processing.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_retry_attempts": {
        "computed": true,
        "description": "The maximum number of times to retry when the function returns an error.",
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "qualifier": {
        "description": "The identifier of a version or alias.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The AWS::Lambda::EventInvokeConfig resource configures options for asynchronous invocation on a version or an alias.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaEventInvokeConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaEventInvokeConfig), &result)
	return &result
}
