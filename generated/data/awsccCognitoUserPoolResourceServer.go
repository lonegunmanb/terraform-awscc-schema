package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolResourceServer = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "scopes": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "scope_description": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "scope_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolResourceServer",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolResourceServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolResourceServer), &result)
	return &result
}
