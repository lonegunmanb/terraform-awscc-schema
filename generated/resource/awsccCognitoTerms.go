package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoTerms = `{
  "block": {
    "attributes": {
      "client_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "enforcement": {
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
      "links": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          "string"
        ]
      },
      "terms_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "terms_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "terms_source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::Terms",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoTermsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoTerms), &result)
	return &result
}
