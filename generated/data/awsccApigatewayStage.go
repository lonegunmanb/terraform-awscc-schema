package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayStage = `{
  "block": {
    "attributes": {
      "access_log_setting": {
        "computed": true,
        "description": "The ` + "`" + `` + "`" + `AccessLogSetting` + "`" + `` + "`" + ` property type specifies settings for logging access in this stage.\n  ` + "`" + `` + "`" + `AccessLogSetting` + "`" + `` + "`" + ` is a property of the [AWS::ApiGateway::Stage](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-stage.html) resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_arn": {
              "computed": true,
              "description": "The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs. If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ` + "`" + `` + "`" + `amazon-apigateway-` + "`" + `` + "`" + `. This parameter is required to enable access logging.",
              "description_kind": "plain",
              "type": "string"
            },
            "format": {
              "computed": true,
              "description": "A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ` + "`" + `` + "`" + `$context.requestId` + "`" + `` + "`" + `. This parameter is required to enable access logging.",
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
        "description_kind": "plain",
        "type": "string"
      },
      "canary_setting": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "deployment_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
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
        "description_kind": "plain",
        "type": "string"
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
      "documentation_version": {
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
      "method_settings": {
        "computed": true,
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
              "description": "The HTTP method. To apply settings to multiple resources and methods, specify an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `) for the ` + "`" + `` + "`" + `HttpMethod` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `/*` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `ResourcePath` + "`" + `` + "`" + `. This parameter is required when you specify a ` + "`" + `` + "`" + `MethodSetting` + "`" + `` + "`" + `.",
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
              "description": "The resource path for this method. Forward slashes (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `) are encoded as ` + "`" + `` + "`" + `~1` + "`" + `` + "`" + ` and the initial slash must include a forward slash. For example, the path value ` + "`" + `` + "`" + `/resource/subresource` + "`" + `` + "`" + ` must be encoded as ` + "`" + `` + "`" + `/~1resource~1subresource` + "`" + `` + "`" + `. To specify the root path, use only a slash (` + "`" + `` + "`" + `/` + "`" + `` + "`" + `). To apply settings to multiple resources and methods, specify an asterisk (` + "`" + `` + "`" + `*` + "`" + `` + "`" + `) for the ` + "`" + `` + "`" + `HttpMethod` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `/*` + "`" + `` + "`" + ` for the ` + "`" + `` + "`" + `ResourcePath` + "`" + `` + "`" + `. This parameter is required when you specify a ` + "`" + `` + "`" + `MethodSetting` + "`" + `` + "`" + `.",
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
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "tracing_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "variables": {
        "computed": true,
        "description": "A map (string-to-string map) that defines the stage variables, where the variable name is the key and the variable value is the value. Variable names are limited to alphanumeric characters. Values must match the following regular expression: ` + "`" + `` + "`" + `[A-Za-z0-9-._~:/?#\u0026=,]+` + "`" + `` + "`" + `.",
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Stage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayStageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayStage), &result)
	return &result
}
