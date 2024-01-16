package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Model = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "The content-type for the model, for example, \"application/json\".",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the model.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "model_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the model.",
        "description_kind": "plain",
        "type": "string"
      },
      "schema": {
        "computed": true,
        "description": "The schema for the model. For application/json models, this should be JSON schema draft 4 model.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::Model",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2ModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Model), &result)
	return &result
}
