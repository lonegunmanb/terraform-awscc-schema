package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2Model = `{
  "block": {
    "attributes": {
      "api_id": {
        "description": "The API identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "content_type": {
        "computed": true,
        "description": "The content-type for the model, for example, \"application/json\".",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the model.",
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
      "model_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the model.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "schema": {
        "description": "The schema for the model. For application/json models, this should be JSON schema draft 4 model.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGatewayV2::Model` + "`" + `` + "`" + ` resource updates data model for a WebSocket API. For more information, see [Model Selection Expressions](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayv2ModelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2Model), &result)
	return &result
}
