package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsChannelAssociation = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "ARN identifier of the channel.\nExample: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_configuration_arn": {
        "computed": true,
        "description": "ARN identifier of the NotificationConfiguration.\nExample: arn:aws:notifications::123456789012:configuration/a01jes88qxwkbj05xv9c967pgm1",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Notifications::ChannelAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNotificationsChannelAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsChannelAssociation), &result)
	return &result
}
