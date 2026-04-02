package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Stage = `{
  "block": {
    "attributes": {
      "access_log_settings": {
        "computed": true,
        "description": "Settings for logging access in this stage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "destination_arn": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "format": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "api_id": {
        "description": "The API identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_deploy": {
        "computed": true,
        "description": "Specifies whether updates to an API automatically trigger a new deployment. The default value is false.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_certificate_id": {
        "computed": true,
        "description": "The identifier of a client certificate for a Stage. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_route_settings": {
        "computed": true,
        "description": "The default route settings for the stage.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "data_trace_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "detailed_metrics_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "logging_level": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "throttling_burst_limit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "throttling_rate_limit": {
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
      "deployment_id": {
        "computed": true,
        "description": "The deployment identifier for the API stage. Can't be updated if autoDeploy is enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description for the API stage.",
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
      "route_settings": {
        "computed": true,
        "description": "Route settings for the stage.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_name": {
        "description": "The stage name. Stage names can contain only alphanumeric characters, hyphens, and underscores, or be $default. Maximum length is 128 characters.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_variables": {
        "computed": true,
        "description": "A map that defines the stage variables for a Stage. Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#\u0026=,]+.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "The collection of tags. Each tag element is associated with a given resource.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::ApiGatewayV2::Stage",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2StageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Stage), &result)
	return &result
}
