package data

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
      "api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_deploy": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "client_certificate_id": {
        "computed": true,
        "description_kind": "plain",
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
              "type": "bool"
            },
            "detailed_metrics_enabled": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "logging_level": {
              "computed": true,
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
      "route_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "stage_variables": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::Stage",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2StageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Stage), &result)
	return &result
}
