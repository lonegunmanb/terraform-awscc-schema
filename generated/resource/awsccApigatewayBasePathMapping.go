package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccApigatewayBasePathMapping = `{
  "block": {
    "attributes": {
      "base_path": {
        "computed": true,
        "description": "The base path name that callers of the API must provide as part of the URL after the domain name.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "The domain name of the BasePathMapping resource to be described.",
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
        "computed": true,
        "description": "The string identifier of the associated RestApi.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "stage": {
        "computed": true,
        "description": "The name of the associated stage.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::ApiGateway::BasePathMapping` + "`" + `` + "`" + ` resource creates a base path that clients who call your API must use in the invocation URL.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccApigatewayBasePathMappingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccApigatewayBasePathMapping), &result)
	return &result
}
