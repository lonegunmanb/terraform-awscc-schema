package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2RouteResponse = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_selection_expression": {
        "computed": true,
        "description": "The model selection expression for the route response. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "type": "string"
      },
      "response_models": {
        "computed": true,
        "description": "The response models for the route response.",
        "description_kind": "plain",
        "type": "string"
      },
      "response_parameters": {
        "computed": true,
        "description": "The route response parameters.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "required": {
              "computed": true,
              "description": "Specifies whether the parameter is required.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "map"
        }
      },
      "route_id": {
        "computed": true,
        "description": "The route ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "route_response_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_response_key": {
        "computed": true,
        "description": "The route response key.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::RouteResponse",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2RouteResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2RouteResponse), &result)
	return &result
}
