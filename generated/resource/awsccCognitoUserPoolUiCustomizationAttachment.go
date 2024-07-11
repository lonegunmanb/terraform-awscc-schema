package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCognitoUserPoolUiCustomizationAttachment = `{
  "block": {
    "attributes": {
      "client_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "css": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      }
    },
    "description": "Resource Type definition for AWS::Cognito::UserPoolUICustomizationAttachment",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCognitoUserPoolUiCustomizationAttachmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCognitoUserPoolUiCustomizationAttachment), &result)
	return &result
}
