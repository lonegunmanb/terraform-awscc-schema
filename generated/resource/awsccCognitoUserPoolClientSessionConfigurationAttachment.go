package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolClientSessionConfigurationAttachment = `{
  "block": {
    "attributes": {
      "client_id": {
        "description": "The ID of the app client.",
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
      "sso_session": {
        "description": "Managed login single sign-on session settings for the app client.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_validity_minutes": {
              "description": "The maximum time in minutes since the last authentication during which this app client can reuse an existing managed login SSO session without asking the user to sign in again. 0 means the app client always asks the user to sign in.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "user_pool_id": {
        "description": "The ID of the user pool.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolClientSessionConfigurationAttachment. Manages the session configuration of a user pool app client. The settings take effect only while the user pool has UserSessionTracking ENABLED, so declaring DependsOn the AWS::Cognito::UserPoolSessionConfigurationAttachment is recommended. The service default (SSOSession MaxValidityMinutes 60) cannot be declared; remove the resource to return the app client to the default.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolClientSessionConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolClientSessionConfigurationAttachment), &result)
	return &result
}
