package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolUserToGroupAttachment = `{
  "block": {
    "attributes": {
      "group_name": {
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
      "user_pool_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "username": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolUserToGroupAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolUserToGroupAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolUserToGroupAttachment), &result)
	return &result
}
