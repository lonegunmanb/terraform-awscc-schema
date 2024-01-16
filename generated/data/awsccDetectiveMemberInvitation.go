package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDetectiveMemberInvitation = `{
  "block": {
    "attributes": {
      "disable_email_notification": {
        "computed": true,
        "description": "When set to true, invitation emails are not sent to the member accounts. Member accounts must still accept the invitation before they are added to the behavior graph. Updating this field has no effect.",
        "description_kind": "plain",
        "type": "bool"
      },
      "graph_arn": {
        "computed": true,
        "description": "The ARN of the graph to which the member account will be invited",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "member_email_address": {
        "computed": true,
        "description": "The root email address for the account to be invited, for validation. Updating this field has no effect.",
        "description_kind": "plain",
        "type": "string"
      },
      "member_id": {
        "computed": true,
        "description": "The AWS account ID to be invited to join the graph as a member",
        "description_kind": "plain",
        "type": "string"
      },
      "message": {
        "computed": true,
        "description": "A message to be included in the email invitation sent to the invited account. Updating this field has no effect.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Detective::MemberInvitation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDetectiveMemberInvitationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDetectiveMemberInvitation), &result)
	return &result
}
