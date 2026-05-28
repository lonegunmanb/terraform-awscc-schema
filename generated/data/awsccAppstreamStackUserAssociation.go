package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamStackUserAssociation = `{
  "block": {
    "attributes": {
      "authentication_type": {
        "computed": true,
        "description": "The authentication type for the user who is associated with the stack. You must specify USERPOOL.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "send_email_notification": {
        "computed": true,
        "description": "Specifies whether a welcome email is sent to a user after the user is created in the user pool.",
        "description_kind": "plain",
        "type": "bool"
      },
      "stack_name": {
        "computed": true,
        "description": "The name of the stack that is associated with the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description": "The name of the user who is associated with the stack.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppStream::StackUserAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamStackUserAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamStackUserAssociation), &result)
	return &result
}
