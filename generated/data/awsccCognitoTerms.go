package data

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
        "type": "string"
      },
      "enforcement": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "links": {
        "computed": true,
        "description_kind": "plain",
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "terms_source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::Terms",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoTermsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoTerms), &result)
	return &result
}
