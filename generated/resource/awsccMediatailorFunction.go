package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccMediatailorFunction = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The ARN of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "custom_output_configuration": {
        "computed": true,
        "description": "Configuration for custom output functions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "output": {
              "computed": true,
              "description": "A map of output key-value pairs that define the custom output.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The runtime environment for the function expression language.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "description": {
        "computed": true,
        "description": "A description of the function.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "function_id": {
        "description": "The unique identifier for the function.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "function_type": {
        "description": "The type of the function. Determines which configuration object is used.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "http_request_configuration": {
        "computed": true,
        "description": "Configuration for HTTP request functions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "body": {
              "computed": true,
              "description": "The body of the HTTP request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "headers": {
              "computed": true,
              "description": "A map of HTTP headers to include in the request.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "method_type": {
              "computed": true,
              "description": "The HTTP method type for the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "output": {
              "computed": true,
              "description": "A map of output key-value pairs. Keys must start with session., temp., avail., scte., or be a valid adsRequest directive.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "request_timeout_milliseconds": {
              "computed": true,
              "description": "The timeout in milliseconds for the HTTP request. Maximum value is 2000.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "runtime": {
              "computed": true,
              "description": "The runtime environment for the function expression language.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "The URL endpoint for the HTTP request.",
              "description_kind": "plain",
              "optional": true,
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
      "sequential_executor_configuration": {
        "computed": true,
        "description": "Configuration for sequential executor functions.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "function_list": {
              "computed": true,
              "description": "The list of functions to execute sequentially.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "function_id": {
                    "computed": true,
                    "description": "The identifier of the function to execute.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "run_condition": {
                    "computed": true,
                    "description": "A conditional expression that determines whether this function should execute.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "output": {
              "computed": true,
              "description": "A map of output key-value pairs that define the final output from sequential execution.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The runtime environment for the function expression language.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "timeout_milliseconds": {
              "computed": true,
              "description": "The timeout in milliseconds for the entire sequential execution chain.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "The tags to assign to the function resource.",
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
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::MediaTailor::Function",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccMediatailorFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorFunction), &result)
	return &result
}
