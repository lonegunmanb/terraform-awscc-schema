package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoIdentityPoolPrincipalTag = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "identity_provider_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "principal_tags": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "use_defaults": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "description": "Data Source schema for AWS::Cognito::IdentityPoolPrincipalTag",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoIdentityPoolPrincipalTagSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoIdentityPoolPrincipalTag), &result)
	return &result
}
