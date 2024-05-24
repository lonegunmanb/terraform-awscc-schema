package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2RouteResponse = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The API identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "model_selection_expression": {
        "computed": true,
        "description": "The model selection expression for the route response. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "response_models": {
        "computed": true,
        "description": "The response models for the route response.",
        "description_kind": "plain",
        "optional": true,
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
              "optional": true,
              "type": "bool"
            }
          },
          "nesting_mode": "map"
        },
        "optional": true
      },
      "route_id": {
        "description": "The route ID.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route_response_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "route_response_key": {
        "description": "The route response key.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGatewayV2::RouteResponse` + "`" + `` + "`" + ` resource creates a route response for a WebSocket API. For more information, see [Set up Route Responses for a WebSocket API in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-route-response.html) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2RouteResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2RouteResponse), &result)
	return &result
}
