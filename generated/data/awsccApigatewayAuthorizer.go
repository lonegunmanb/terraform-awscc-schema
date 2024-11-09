package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayAuthorizer = `{
  "block": {
    "attributes": {
      "auth_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_credentials": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "authorizer_result_ttl_in_seconds": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "authorizer_uri": {
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
      "identity_source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_validation_expression": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provider_ar_ns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "rest_api_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::Authorizer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayAuthorizerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayAuthorizer), &result)
	return &result
}
