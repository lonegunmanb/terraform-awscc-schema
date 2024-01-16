package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLambdaUrl = `{
  "block": {
    "attributes": {
      "auth_type": {
        "description": "Can be either AWS_IAM if the requests are authorized via IAM, or NONE if no authorization is configured on the Function URL.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cors": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "allow_credentials": {
              "computed": true,
              "description": "Specifies whether credentials are included in the CORS request.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "allow_headers": {
              "computed": true,
              "description": "Represents a collection of allowed headers.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_methods": {
              "computed": true,
              "description": "Represents a collection of allowed HTTP methods.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "allow_origins": {
              "computed": true,
              "description": "Represents a collection of allowed origins.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "expose_headers": {
              "computed": true,
              "description": "Represents a collection of exposed headers.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "max_age": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "function_arn": {
        "computed": true,
        "description": "The full Amazon Resource Name (ARN) of the function associated with the Function URL.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_url": {
        "computed": true,
        "description": "The generated url for this resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "invoke_mode": {
        "computed": true,
        "description": "The invocation mode for the function?s URL. Set to BUFFERED if you want to buffer responses before returning them to the client. Set to RESPONSE_STREAM if you want to stream responses, allowing faster time to first byte and larger response payload sizes. If not set, defaults to BUFFERED.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "qualifier": {
        "computed": true,
        "description": "The alias qualifier for the target function. If TargetFunctionArn is unqualified then Qualifier must be passed.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "target_function_arn": {
        "description": "The Amazon Resource Name (ARN) of the function associated with the Function URL.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Lambda::Url",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLambdaUrlSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLambdaUrl), &result)
	return &result
}
