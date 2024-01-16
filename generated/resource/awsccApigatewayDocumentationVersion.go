package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayDocumentationVersion = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A description about the new documentation snapshot.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "documentation_version": {
        "description": "The version identifier of the to-be-updated documentation version.",
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
      "rest_api_id": {
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::DocumentationVersion` + "`" + `` + "`" + ` resource creates a snapshot of the documentation for an API. For more information, see [Representation of API Documentation in API Gateway](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-documenting-api-content-representation.html) in the *API Gateway Developer Guide*.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayDocumentationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDocumentationVersion), &result)
	return &result
}
