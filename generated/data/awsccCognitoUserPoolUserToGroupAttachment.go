package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolUserToGroupAttachment = `{
  "block": {
    "attributes": {
      "group_name": {
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
      "user_pool_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Cognito::UserPoolUserToGroupAttachment",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCognitoUserPoolUserToGroupAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolUserToGroupAttachment), &result)
	return &result
}
