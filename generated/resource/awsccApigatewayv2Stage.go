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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_deploy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_certificate_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "default_route_settings": {
        "computed": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
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
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage_variables": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
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
