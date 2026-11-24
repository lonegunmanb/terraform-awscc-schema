package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolSessionConfigurationAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "token_session": {
        "description": "Token session settings that apply to every app client in the user pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_per_user": {
              "description": "The maximum number of concurrent token sessions per user across all app clients. When a new session would exceed this number, the least recently used session ends.",
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
      },
      "user_session_tracking": {
        "description": "Whether server-side user session tracking is enabled for the user pool. Requires the ESSENTIALS or PLUS feature plan.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolSessionConfigurationAttachment. Manages the session configuration of a user pool: whether user session tracking is enabled and the maximum number of concurrent token sessions per user. The service default (UserSessionTracking DISABLED with TokenSession MaxPerUser 20) cannot be declared; remove the resource to return the user pool to the default.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolSessionConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolSessionConfigurationAttachment), &result)
	return &result
}
