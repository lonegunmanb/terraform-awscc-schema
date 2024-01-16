package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2ApiMapping = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The identifier of the API.",
        "description_kind": "plain",
        "type": "string"
      },
      "api_mapping_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "api_mapping_key": {
        "computed": true,
        "description": "The API mapping key.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description": "The domain name.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description": "The API stage.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::ApiMapping",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2ApiMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2ApiMapping), &result)
	return &result
}
