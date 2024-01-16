package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayv2IntegrationResponse = `{
  "block": {
    "attributes": {
      "api_id": {
        "computed": true,
        "description": "The API identifier.",
        "description_kind": "plain",
        "type": "string"
      },
      "content_handling_strategy": {
        "computed": true,
        "description": "Supported only for WebSocket APIs. Specifies how to handle response payload content type conversions. Supported values are ` + "`" + `` + "`" + `CONVERT_TO_BINARY` + "`" + `` + "`" + ` and ` + "`" + `` + "`" + `CONVERT_TO_TEXT` + "`" + `` + "`" + `, with the following behaviors:\n  ` + "`" + `` + "`" + `CONVERT_TO_BINARY` + "`" + `` + "`" + `: Converts a response payload from a Base64-encoded string to the corresponding binary blob.\n  ` + "`" + `` + "`" + `CONVERT_TO_TEXT` + "`" + `` + "`" + `: Converts a response payload from a binary blob to a Base64-encoded string.\n If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "integration_id": {
        "computed": true,
        "description": "The integration ID.",
        "description_kind": "plain",
        "type": "string"
      },
      "integration_response_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "integration_response_key": {
        "computed": true,
        "description": "The integration response key.",
        "description_kind": "plain",
        "type": "string"
      },
      "response_parameters": {
        "computed": true,
        "description": "A key-value map specifying response parameters that are passed to the method response from the backend. The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ` + "`" + `` + "`" + `method.response.header.{name}` + "`" + `` + "`" + `, where name is a valid and unique header name. The mapped non-static value must match the pattern of ` + "`" + `` + "`" + `integration.response.header.{name}` + "`" + `` + "`" + ` or ` + "`" + `` + "`" + `integration.response.body.{JSON-expression}` + "`" + `` + "`" + `, where ` + "`" + `` + "`" + `{name}` + "`" + `` + "`" + ` is a valid and unique response header name and ` + "`" + `` + "`" + `{JSON-expression}` + "`" + `` + "`" + ` is a valid JSON expression without the ` + "`" + `` + "`" + `$` + "`" + `` + "`" + ` prefix.",
        "description_kind": "plain",
        "type": "string"
      },
      "response_templates": {
        "computed": true,
        "description": "The collection of response templates for the integration response as a string-to-string map of key-value pairs. Response templates are represented as a key/value map, with a content-type as the key and a template as the value.",
        "description_kind": "plain",
        "type": "string"
      },
      "template_selection_expression": {
        "computed": true,
        "description": "The template selection expression for the integration response. Supported only for WebSocket APIs.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGatewayV2::IntegrationResponse",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayv2IntegrationResponseSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayv2IntegrationResponse), &result)
	return &result
}
