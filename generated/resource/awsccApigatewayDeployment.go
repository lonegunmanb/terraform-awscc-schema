package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDeployment = `{
  "block": {
    "attributes": {
      "deployment_canary_settings": {
        "computed": true,
        "description": "The input configuration for a canary deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "percent_traffic": {
              "computed": true,
              "description": "The percentage (0.0-100.0) of traffic routed to the canary deployment.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "stage_variable_overrides": {
              "computed": true,
              "description": "A stage variable overrides used for the canary release deployment. They can override existing stage variables or add new stage variables for the canary release deployment. These stage variables are represented as a string-to-string map between stage variable names and their values.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "use_stage_cache": {
              "computed": true,
              "description": "A Boolean flag to indicate whether the canary release deployment uses the stage cache or not.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for the Deployment resource to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_description": {
        "computed": true,
        "description": "The description of the Stage resource for the Deployment resource to create. To specify a stage description, you must also provide a stage name.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "access_log_setting": {
              "computed": true,
              "description": "Specifies settings for logging access in this stage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "destination_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ` + "`" + `` + "`" + `amazon-apigateway-` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "format": {
                    "computed": true,
                    "description": "A single line format of the access logs of data, as specified by selected $context variables. The format must include at least ` + "`" + `` + "`" + `$context.requestId` + "`" + `` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "cache_cluster_enabled": {
              "computed": true,
              "description": "Specifies whether a cache cluster is enabled for the stage.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cache_cluster_size": {
              "computed": true,
              "description": "The size of the stage's cache cluster. For more information, see [cacheClusterSize](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html#apigw-CreateStage-request-cacheClusterSize) in the *API Gateway API Reference*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cache_data_encrypted": {
              "computed": true,
              "description": "Indicates whether the cached responses are encrypted.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "cache_ttl_in_seconds": {
              "computed": true,
              "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "caching_enabled": {
              "computed": true,
              "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "canary_setting": {
              "computed": true,
              "description": "Specifies settings for the canary deployment in this stage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "percent_traffic": {
                    "computed": true,
                    "description": "The percent (0-100) of traffic diverted to a canary deployment.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "stage_variable_overrides": {
                    "computed": true,
                    "description": "Stage variables overridden for a canary release deployment, including new stage variables introduced in the canary. These stage variables are represented as a string-to-string map between stage variable names and their values.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "use_stage_cache": {
                    "computed": true,
                    "description": "A Boolean flag to indicate whether the canary deployment uses the stage cache or not.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "client_certificate_id": {
              "computed": true,
              "description": "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "data_trace_enabled": {
              "computed": true,
              "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "description": {
              "computed": true,
              "description": "A description of the purpose of the stage.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "documentation_version": {
              "computed": true,
              "description": "The version identifier of the API documentation snapshot.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "logging_level": {
              "computed": true,
              "description": "The logging level for this method. For valid values, see the ` + "`" + `` + "`" + `loggingLevel` + "`" + `` + "`" + ` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the *Amazon API Gateway API Reference*.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "method_settings": {
              "computed": true,
              "description": "Configures settings for all of the stage's methods.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "cache_data_encrypted": {
                    "computed": true,
                    "description": "Specifies whether the cached responses are encrypted.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "cache_ttl_in_seconds": {
                    "computed": true,
                    "description": "Specifies the time to live (TTL), in seconds, for cached responses. The higher the TTL, the longer the response will be cached.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "caching_enabled": {
                    "computed": true,
                    "description": "Specifies whether responses should be cached and returned for requests. A cache cluster must be enabled on the stage for responses to be cached.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "data_trace_enabled": {
                    "computed": true,
                    "description": "Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs. This can be useful to troubleshoot APIs, but can result in logging sensitive data. We recommend that you don't enable this option for production APIs.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "http_method": {
                    "computed": true,
                    "description": "The HTTP method.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "logging_level": {
                    "computed": true,
                    "description": "Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs. Valid values are ` + "`" + `` + "`" + `OFF` + "`" + `` + "`" + `, ` + "`" + `` + "`" + `ERROR` + "`" + `` + "`" + `, and ` + "`" + `` + "`" + `INFO` + "`" + `` + "`" + `. Choose ` + "`" + `` + "`" + `ERROR` + "`" + `` + "`" + ` to write only error-level entries to CloudWatch Logs, or choose ` + "`" + `` + "`" + `INFO` + "`" + `` + "`" + ` to include all ` + "`" + `` + "`" + `ERROR` + "`" + `` + "`" + ` events as well as extra informational events.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metrics_enabled": {
                    "computed": true,
                    "description": "Specifies whether Amazon CloudWatch metrics are enabled for this method.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "resource_path": {
                    "computed": true,
                    "description": "The resource path for this method. Forward slashes (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) are encoded as ` + "`" + `` + "`" + `~1` + "`" + `` + "`" + ` and the initial slash must include a forward slash. For example, the path value ` + "`" + `` + "`" + `/resource/subresource` + "`" + `` + "`" + ` must be encoded as ` + "`" + `` + "`" + `/~1resource~1subresource` + "`" + `` + "`" + `. To specify the root path, use only a slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `).",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "throttling_burst_limit": {
                    "computed": true,
                    "description": "Specifies the throttling burst limit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "throttling_rate_limit": {
                    "computed": true,
                    "description": "Specifies the throttling rate limit.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              },
              "optional": true
            },
            "metrics_enabled": {
              "computed": true,
              "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tags": {
              "computed": true,
              "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "description": "The key name of the tag",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "value": {
                    "description": "The value for the tag",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              },
              "optional": true
            },
            "throttling_burst_limit": {
              "computed": true,
              "description": "The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throttling_rate_limit": {
              "computed": true,
              "description": "The target request steady-state rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tracing_enabled": {
              "computed": true,
              "description": "Specifies whether active tracing with X-ray is enabled for this stage.\n For more information, see [Trace API Gateway API Execution with X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "variables": {
              "computed": true,
              "description": "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: ` + "`" + `` + "`" + `[A-Za-z0-9-._~:/?#\u0026=,]+` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "stage_name": {
        "computed": true,
        "description": "The name of the Stage resource for the Deployment resource to create.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Deployment` + "`" + `` + "`" + ` resource deploys an API Gateway ` + "`" + `` + "`" + `RestApi` + "`" + `` + "`" + ` resource to a stage so that clients can call the API over the internet. The stage acts as an environment.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDeployment), &result)
	return &result
}
