package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccAppstreamUser = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "Returns the Amazon Resource Name (ARN) for the Amazon AppStream User resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "authentication_type": {
        "computed": true,
        "description": "The authentication type for the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "first_name": {
        "computed": true,
        "description": "The first name, or given name, of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "last_name": {
        "computed": true,
        "description": "The last name, or surname, of the user.",
        "description_kind": "plain",
        "type": "string"
      },
      "message_action": {
        "computed": true,
        "description": "The action to take for the welcome email that is sent to a user after the user is created in the user pool. If you specify SUPPRESS, no email is sent. If you specify RESEND, do not specify the first name or last name of the user. If the value is null, the email is sent.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description": "The email address of the user.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::AppStream::User",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccAppstreamUserSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccAppstreamUser), &result)
	return &result
}
