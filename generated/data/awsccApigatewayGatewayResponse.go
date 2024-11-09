package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayGatewayResponse = `{
  "block": {
    "attributes": {
      "gateway_response_id": {
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
      "response_parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "response_templates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "response_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_code": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::GatewayResponse",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayGatewayResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayGatewayResponse), &result)
	return &result
}
