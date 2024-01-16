package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayStage = `{
  "block": {
    "attributes": {
      "access_log_setting": {
        "computed": true,
        "description": "Access log settings, including the access log format and access log destination ARN.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ` + "`" + `` + "`" + `amazon-apigateway-` + "`" + `` + "`" + `. This parameter is required to enable access logging.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format": {
              "computed": true,
              "description": "A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ` + "`" + `` + "`" + `$context.requestId` + "`" + `` + "`" + `. This parameter is required to enable access logging.",
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
        "description": "The stage's cache capacity in GB. For more information about choosing a cache size, see [Enabling API caching to enhance responsiveness](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-caching.html).",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "canary_setting": {
        "computed": true,
        "description": "Settings for the canary deployment in this stage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_id": {
              "computed": true,
              "description": "The ID of the canary deployment.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
        "description": "The identifier of a client certificate for an API stage.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "deployment_id": {
        "computed": true,
        "description": "The identifier of the Deployment that the stage points to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The stage's description.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "documentation_version": {
        "computed": true,
        "description": "The version of the associated API documentation.",
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
      "method_settings": {
        "computed": true,
        "description": "A map that defines the method settings for a Stage resource. Keys (designated as ` + "`" + `` + "`" + `/{method_setting_key` + "`" + `` + "`" + ` below) are method paths defined as ` + "`" + `` + "`" + `{resource_path}/{http_method}` + "`" + `` + "`" + ` for an individual method override, or ` + "`" + `` + "`" + `/\\*/\\*` + "`" + `` + "`" + ` for overriding all methods in the stage.",
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
              "description": "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `) for the ` + "`" + `` + "`" + `HttpMethod` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `/*` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `ResourcePath` + "`" + `` + "`" + `. This parameter is required when you specify a ` + "`" + `` + "`" + `MethodSetting` + "`" + `` + "`" + `.",
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
              "description": "The resource path for this method. Forward slashes (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) are encoded as ` + "`" + `` + "`" + `~1` + "`" + `` + "`" + ` and the initial slash must include a forward slash. For example, the path value ` + "`" + `` + "`" + `/resource/subresource` + "`" + `` + "`" + ` must be encoded as ` + "`" + `` + "`" + `/~1resource~1subresource` + "`" + `` + "`" + `. To specify the root path, use only a slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `). To apply settings to multiple resources and methods, specify an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `) for the ` + "`" + `` + "`" + `HttpMethod` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `/*` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `ResourcePath` + "`" + `` + "`" + `. This parameter is required when you specify a ` + "`" + `` + "`" + `MethodSetting` + "`" + `` + "`" + `.",
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
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_name": {
        "computed": true,
        "description": "The name of the stage is the first path segment in the Uniform Resource Identifier (URI) of a call to API Gateway. Stage names can only contain alphanumeric characters, hyphens, and underscores. Maximum length is 128 characters.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "tracing_enabled": {
        "computed": true,
        "description": "Specifies whether active tracing with X-ray is enabled for the Stage.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "variables": {
        "computed": true,
        "description": "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ` + "`" + `` + "`" + `[A-Za-z0-9-._~:/?#\u0026=,]+` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Stage` + "`" + `` + "`" + ` resource creates a stage for a deployment.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayStageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayStage), &result)
	return &result
}
