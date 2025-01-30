package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsManagedNotificationAdditionalChannelAssociation = `{
  "block": {
    "attributes": {
      "channel_arn": {
        "description": "ARN identifier of the channel.\nExample: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops",
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
      "managed_notification_configuration_arn": {
        "description": "ARN identifier of the Managed Notification.\nExample: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::Notifications::ManagedNotificationAdditionalChannelAssociation Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNotificationsManagedNotificationAdditionalChannelAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsManagedNotificationAdditionalChannelAssociation), &result)
	return &result
}
