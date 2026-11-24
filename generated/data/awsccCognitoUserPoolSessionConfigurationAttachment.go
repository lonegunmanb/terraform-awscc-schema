package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolSessionConfigurationAttachment = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "token_session": {
        "computed": true,
        "description": "Token session settings that apply to every app client in the user pool.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "max_per_user": {
              "computed": true,
              "description": "The maximum number of concurrent token sessions per user across all app clients. When a new session would exceed this number, the least recently used session ends.",
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
      },
      "user_session_tracking": {
        "computed": true,
        "description": "Whether server-side user session tracking is enabled for the user pool. Requires the ESSENTIALS or PLUS feature plan.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolSessionConfigurationAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolSessionConfigurationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolSessionConfigurationAttachment), &result)
	return &result
}
