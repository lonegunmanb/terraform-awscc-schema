package data

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
        "type": "string"
      },
      "documentation_version": {
        "computed": true,
        "description": "The version identifier of the to-be-updated documentation version.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rest_api_id": {
        "computed": true,
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ApiGateway::DocumentationVersion",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccApigatewayDocumentationVersionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayDocumentationVersion), &result)
	return &result
}
