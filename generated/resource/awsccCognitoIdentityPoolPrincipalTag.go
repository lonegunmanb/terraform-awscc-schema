package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoIdentityPoolPrincipalTag = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_provider_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "principal_tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_defaults": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::IdentityPoolPrincipalTag",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoIdentityPoolPrincipalTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoIdentityPoolPrincipalTag), &result)
	return &result
}
