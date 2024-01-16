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
        "description": "The name of this RequestValidator",
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
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "type": "string"
      },
      "validate_request_body": {
        "computed": true,
        "description": "A Boolean flag to indicate whether to validate a request body according to the configured Model schema.",
        "description_kind": "plain",
        "type": "bool"
      },
      "validate_request_parameters": {
        "computed": true,
        "description": "A Boolean flag to indicate whether to validate request parameters (` + "`" + `` + "`" + `true` + "`" + `` + "`" + `) or not (` + "`" + `` + "`" + `false` + "`" + `` + "`" + `).",
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
