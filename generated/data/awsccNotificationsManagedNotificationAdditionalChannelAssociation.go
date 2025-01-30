package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsManagedNotificationAdditionalChannelAssociation = `{
  "block": {
    "attributes": {
      "channel_arn": {
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
      "managed_notification_configuration_arn": {
        "computed": true,
        "description": "ARN identifier of the Managed Notification.\nExample: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Notifications::ManagedNotificationAdditionalChannelAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNotificationsManagedNotificationAdditionalChannelAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsManagedNotificationAdditionalChannelAssociation), &result)
	return &result
}
