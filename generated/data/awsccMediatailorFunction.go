package data

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
      "aws_service_request_configuration": {
        "computed": true,
        "description": "The configuration for an AWS_SERVICE_REQUEST function. Contains the target service, target Region, and request parameters that the function uses to call an AWS service API. For more information, see AWS_SERVICE_REQUEST (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types-aws-service-request.html) in the MediaTailor User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "body": {
              "computed": true,
              "description": "An expression that evaluates to the request body for the AWS service API call. The body must conform to the input format that the target service operation expects. Applies only when the target operation accepts a request body. The maximum size after evaluation is 64 KB.",
              "description_kind": "plain",
              "type": "string"
            },
            "headers": {
              "computed": true,
              "description": "A map of HTTP header names to expression values. MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request to the AWS service. Use this to pass any headers required by the target service operation. You can include a maximum of 50 headers.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "method_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "output": {
              "computed": true,
              "description": "A map of output bindings. Each key is a namespaced output path, such as player_params.device_type. Each value is an expression that MediaTailor evaluates at runtime and can reference the response object from the target service. For more information, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "request_timeout_milliseconds": {
              "computed": true,
              "description": "The maximum time, in milliseconds, that MediaTailor waits for a response from the AWS service. If the call exceeds this timeout, MediaTailor sets the response status code to null and proceeds with output expression evaluation. Valid values are 100 to 2000.",
              "description_kind": "plain",
              "type": "number"
            },
            "runtime": {
              "computed": true,
              "description": "The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_region": {
              "computed": true,
              "description": "The AWS Region for the target service. Specify a static Region code (for example, us-east-1) or a JSONata expression that resolves to a Region code at runtime (for example, {%inference.region%}).",
              "description_kind": "plain",
              "type": "string"
            },
            "target_service": {
              "computed": true,
              "description": "The AWS service to call. Valid value: elemental-inference (AWS Elemental Inference).",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "An expression that evaluates to the endpoint URL for the target AWS service API operation. Use {%...%} delimiters for dynamic expressions. The URL must correspond to a valid endpoint for the service specified in TargetService. The maximum length after evaluation is 2,048 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "concurrent_executor_configuration": {
        "computed": true,
        "description": "The configuration for a CONCURRENT_EXECUTOR function. Required when FunctionType is CONCURRENT_EXECUTOR.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "function_list": {
              "computed": true,
              "description": "The list of 1 to 10 child functions that MediaTailor runs in parallel. Each entry specifies a child function to execute and an optional run condition expression that controls whether the function runs. Child functions cannot themselves be executors, and each child function's resolved namespace must be unique across the list.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alias": {
                    "computed": true,
                    "description": "An optional alternate name for the child function within the executor. MediaTailor uses this value as the namespace for the child function's output. If omitted, MediaTailor uses the function identifier. The resolved namespace must be unique across all child functions in the list.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "function_id": {
                    "computed": true,
                    "description": "The identifier of the child function to execute.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "run_condition": {
                    "computed": true,
                    "description": "An optional expression that evaluates to a boolean. MediaTailor evaluates this expression immediately before running the child function, using the accumulated state at that point. If the expression evaluates to false, MediaTailor skips the child function. If omitted, the child function always runs.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "max_concurrency": {
              "computed": true,
              "description": "The maximum number of child functions that MediaTailor runs simultaneously. When the list contains more functions than MaxConcurrency, MediaTailor starts additional functions as running ones complete, so that no more than MaxConcurrency functions run at the same time. Valid values are 1 to 2.",
              "description_kind": "plain",
              "type": "number"
            },
            "output": {
              "computed": true,
              "description": "A map of output bindings that controls which bindings the executor commits to the session state after all child functions complete. Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the combined results of the child functions.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.",
              "description_kind": "plain",
              "type": "string"
            },
            "timeout_milliseconds": {
              "computed": true,
              "description": "The maximum time, in milliseconds, for all child functions to complete. This timeout covers every function in the list, including any HTTP calls the child functions make. If the executor exceeds this timeout, MediaTailor discards all output from the executor and proceeds with default behavior. Valid values are 100 to 2000.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The runtime environment for the function expression language.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "description": {
        "computed": true,
        "description": "A description of the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_id": {
        "computed": true,
        "description": "The unique identifier for the function.",
        "description_kind": "plain",
        "type": "string"
      },
      "function_type": {
        "computed": true,
        "description": "The type of the function. Determines which configuration object is used.",
        "description_kind": "plain",
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
              "type": "string"
            },
            "headers": {
              "computed": true,
              "description": "A map of HTTP headers to include in the request.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "method_type": {
              "computed": true,
              "description": "The HTTP method type for the request.",
              "description_kind": "plain",
              "type": "string"
            },
            "output": {
              "computed": true,
              "description": "A map of output key-value pairs. Keys must start with session., temp., avail., scte., or be a valid adsRequest directive.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "request_timeout_milliseconds": {
              "computed": true,
              "description": "The timeout in milliseconds for the HTTP request. Maximum value is 2000.",
              "description_kind": "plain",
              "type": "number"
            },
            "runtime": {
              "computed": true,
              "description": "The runtime environment for the function expression language.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "The URL endpoint for the HTTP request.",
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
      "sequential_executor_configuration": {
        "computed": true,
        "description": "The configuration for a SEQUENTIAL_EXECUTOR function. A SEQUENTIAL_EXECUTOR runs an ordered list of child functions one at a time, passing data between them. For more information about functions, see Working with functions (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions.html) in the MediaTailor User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "function_list": {
              "computed": true,
              "description": "An ordered list of 1 to 10 steps. Each step specifies a child function to execute and an optional run condition expression that controls whether the step runs. MediaTailor executes the steps in order, passing data between steps through temporary data. Each step's resolved namespace must be unique across the list.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "alias": {
                    "computed": true,
                    "description": "An optional alternate name for the child function within the executor. MediaTailor uses this value as the namespace for the child function's output. If omitted, MediaTailor uses the function identifier. The resolved namespace must be unique across all child functions in the list.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "function_id": {
                    "computed": true,
                    "description": "The identifier of the child function to execute.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "run_condition": {
                    "computed": true,
                    "description": "An optional expression that evaluates to a boolean. MediaTailor evaluates this expression immediately before running the child function, using the accumulated state at that point. If the expression evaluates to false, MediaTailor skips the child function. If omitted, the child function always runs.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "output": {
              "computed": true,
              "description": "A map of output bindings that controls which bindings the sequence commits to the session state after all steps complete. Each key is a namespaced output path, and each value is an expression that MediaTailor evaluates against the accumulated results of the steps.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "runtime": {
              "computed": true,
              "description": "The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.",
              "description_kind": "plain",
              "type": "string"
            },
            "timeout_milliseconds": {
              "computed": true,
              "description": "The maximum time, in milliseconds, for the entire sequence to complete. This timeout covers all steps, including any HTTP calls made by child functions. If the sequence exceeds this timeout, MediaTailor discards all output from the sequence and proceeds with default behavior. Valid values are 100 to 2000.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
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
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      },
      "vast_request_configuration": {
        "computed": true,
        "description": "The configuration for a VAST_REQUEST function. Specifies the HTTP method, URL, headers, body, timeout, and output expressions for a request to a VAST endpoint. MediaTailor parses the response as VAST and resolves wrapper redirects, then makes the parsed ads available to the function's output expressions. For more information, see Function types and composition (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-types.html) in the MediaTailor User Guide.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "body": {
              "computed": true,
              "description": "An expression that evaluates to the request body, for example to send an OpenRTB bid request. The expression can be up to 100,000 characters, and the body after evaluation can be up to 64 KB.",
              "description_kind": "plain",
              "type": "string"
            },
            "headers": {
              "computed": true,
              "description": "A map of HTTP header names to expression values. MediaTailor evaluates each header value expression at runtime and includes the result in the outbound request. Headers beginning with X-Amz- are reserved by the service, and method override headers are not allowed.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "method_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "output": {
              "computed": true,
              "description": "A map of output bindings. Each key is a namespaced output path (such as temp.wrappedAds), and each value is an expression that MediaTailor evaluates at runtime. Output expressions in a VAST_REQUEST function can reference the response object, which exposes response.parsedAds, the ads parsed from the VAST response after schema validation and wrapper resolution, and response.statusCode. For more information about expression syntax, see JSONata expression reference (https://docs.aws.amazon.com/mediatailor/latest/ug/monetization-functions-jsonata.html) in the MediaTailor User Guide.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "request_timeout_milliseconds": {
              "computed": true,
              "description": "The maximum time, in milliseconds, that MediaTailor waits for a response from the VAST endpoint. The timeout covers the entire response, including any wrapper redirects that MediaTailor follows. If the call exceeds this timeout, MediaTailor proceeds with an empty ad list and continues output expression evaluation. Valid values are 100 to 2000.",
              "description_kind": "plain",
              "type": "number"
            },
            "runtime": {
              "computed": true,
              "description": "The expression language used to evaluate expressions in the function configuration. Set this to JSONATA.",
              "description_kind": "plain",
              "type": "string"
            },
            "url": {
              "computed": true,
              "description": "An expression that evaluates to the VAST endpoint URL. Use {%...%} delimiters for dynamic expressions. A literal value must be an https:// URL. The expression can be up to 25,000 characters, and the URL after evaluation can be up to 2,048 characters.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::MediaTailor::Function",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccMediatailorFunctionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccMediatailorFunction), &result)
	return &result
}
