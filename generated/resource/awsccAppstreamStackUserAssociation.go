package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamStackUserAssociation = `{
  "block": {
    "attributes": {
      "authentication_type": {
        "description": "The authentication type for the user who is associated with the stack. You must specify USERPOOL.",
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
      "send_email_notification": {
        "computed": true,
        "description": "Specifies whether a welcome email is sent to a user after the user is created in the user pool.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "stack_name": {
        "description": "The name of the stack that is associated with the user.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "description": "The name of the user who is associated with the stack.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::AppStream::StackUserAssociation",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccAppstreamStackUserAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamStackUserAssociation), &result)
	return &result
}
