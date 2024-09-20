package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolResourceServer = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
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
              "optional": true,
              "type": "string"
            },
            "scope_name": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolResourceServer",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolResourceServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolResourceServer), &result)
	return &result
}
