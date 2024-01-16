package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayResource = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_id": {
        "description": "The parent resource's identifier.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path_part": {
        "description": "The last path segment for this resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "computed": true,
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
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::Resource` + "`" + `` + "`" + ` resource creates a resource in an API.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayResource), &result)
	return &result
}
