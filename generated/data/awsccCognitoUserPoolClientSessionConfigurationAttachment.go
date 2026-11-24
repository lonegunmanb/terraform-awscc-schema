package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolClientSessionConfigurationAttachment = `{
  "block": {
    "attributes": {
      "client_id": {
        "computed": true,
        "description": "The ID of the app client.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sso_session": {
        "computed": true,
        "description": "Managed login single sign-on session settings for the app client.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_validity_minutes": {
              "computed": true,
              "description": "The maximum time in minutes since the last authentication during which this app client can reuse an existing managed login SSO session without asking the user to sign in again. 0 means the app client always asks the user to sign in.",
              "description_kind": "plain",
              "type": "number"
            }
          },
          "nesting_mode": "single"
        }
      },
      "user_pool_id": {
        "computed": true,
        "description": "The ID of the user pool.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolClientSessionConfigurationAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolClientSessionConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolClientSessionConfigurationAttachment), &result)
	return &result
}
