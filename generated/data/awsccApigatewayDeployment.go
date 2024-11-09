package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDeployment = `{
  "block": {
    "attributes": {
      "deployment_canary_settings": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `DeploymentCanarySettings` + "`" + `` + "`" + ` property type specifies settings for the canary deployment.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "percent_traffic": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "stage_variable_overrides": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            },
            "use_stage_cache": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
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
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "format": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "cache_cluster_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "cache_cluster_size": {
              "computed": true,
              "description": "The size of the stage's cache cluster. For more information, see [cacheClusterSize](https://docs.aws.amazon.com/apigateway/latest/api/API_CreateStage.html#apigw-CreateStage-request-cacheClusterSize) in the *API Gateway API Reference*.",
              "description_kind": "plain",
              "type": "string"
            },
            "cache_data_encrypted": {
              "computed": true,
              "description": "Indicates whether the cached responses are encrypted.",
              "description_kind": "plain",
              "type": "bool"
            },
            "cache_ttl_in_seconds": {
              "computed": true,
              "description": "The time-to-live (TTL) period, in seconds, that specifies how long API Gateway caches responses.",
              "description_kind": "plain",
              "type": "number"
            },
            "caching_enabled": {
              "computed": true,
              "description": "Indicates whether responses are cached and returned for requests. You must enable a cache cluster on the stage to cache responses. For more information, see [Enable API Gateway Caching in a Stage to Enhance API Performance](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
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
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "stage_variable_overrides": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "use_stage_cache": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "client_certificate_id": {
              "computed": true,
              "description": "The identifier of the client certificate that API Gateway uses to call your integration endpoints in the stage.",
              "description_kind": "plain",
              "type": "string"
            },
            "data_trace_enabled": {
              "computed": true,
              "description": "Indicates whether data trace logging is enabled for methods in the stage. API Gateway pushes these logs to Amazon CloudWatch Logs.",
              "description_kind": "plain",
              "type": "bool"
            },
            "description": {
              "computed": true,
              "description": "A description of the purpose of the stage.",
              "description_kind": "plain",
              "type": "string"
            },
            "documentation_version": {
              "computed": true,
              "description": "The version identifier of the API documentation snapshot.",
              "description_kind": "plain",
              "type": "string"
            },
            "logging_level": {
              "computed": true,
              "description": "The logging level for this method. For valid values, see the ` + "`" + `` + "`" + `loggingLevel` + "`" + `` + "`" + ` property of the [MethodSetting](https://docs.aws.amazon.com/apigateway/latest/api/API_MethodSetting.html) resource in the *Amazon API Gateway API Reference*.",
              "description_kind": "plain",
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
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "cache_ttl_in_seconds": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "caching_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "data_trace_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "http_method": {
                    "computed": true,
                    "description": "The HTTP method.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "logging_level": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "metrics_enabled": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "bool"
                  },
                  "resource_path": {
                    "computed": true,
                    "description": "The resource path for this method. Forward slashes (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) are encoded as ` + "`" + `` + "`" + `~1` + "`" + `` + "`" + ` and the initial slash must include a forward slash. For example, the path value ` + "`" + `` + "`" + `/resource/subresource` + "`" + `` + "`" + ` must be encoded as ` + "`" + `` + "`" + `/~1resource~1subresource` + "`" + `` + "`" + `. To specify the root path, use only a slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `).",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "throttling_burst_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  },
                  "throttling_rate_limit": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "number"
                  }
                },
                "nesting_mode": "set"
              }
            },
            "metrics_enabled": {
              "computed": true,
              "description": "Indicates whether Amazon CloudWatch metrics are enabled for methods in the stage.",
              "description_kind": "plain",
              "type": "bool"
            },
            "tags": {
              "computed": true,
              "description": "An array of arbitrary tags (key-value pairs) to associate with the stage.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "key": {
                    "computed": true,
                    "description": "The key name of the tag",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The value for the tag",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "list"
              }
            },
            "throttling_burst_limit": {
              "computed": true,
              "description": "The target request burst rate limit. This allows more requests through for a period of time than the target rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "type": "number"
            },
            "throttling_rate_limit": {
              "computed": true,
              "description": "The target request steady-state rate limit. For more information, see [Manage API Request Throttling](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "type": "number"
            },
            "tracing_enabled": {
              "computed": true,
              "description": "Specifies whether active tracing with X-ray is enabled for this stage.\n For more information, see [Trace API Gateway API Execution with X-Ray](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-xray.html) in the *API Gateway Developer Guide*.",
              "description_kind": "plain",
              "type": "bool"
            },
            "variables": {
              "computed": true,
              "description": "A map that defines the stage variables. Variable names must consist of alphanumeric characters, and the values must match the following regular expression: ` + "`" + `` + "`" + `[A-Za-z0-9-._~:/?#\u0026=,]+` + "`" + `` + "`" + `.",
              "description_kind": "plain",
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "single"
        }
      },
      "stage_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Deployment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDeployment), &result)
	return &result
}
