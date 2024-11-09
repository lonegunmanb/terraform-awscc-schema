package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayRequestValidator = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "request_validator_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "validate_request_body": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "validate_request_parameters": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::RequestValidator",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayRequestValidatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayRequestValidator), &result)
	return &result
}
