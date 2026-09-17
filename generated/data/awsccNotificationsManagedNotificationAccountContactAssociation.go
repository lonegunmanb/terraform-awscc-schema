package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsManagedNotificationAccountContactAssociation = `{
  "block": {
    "attributes": {
      "contact_identifier": {
        "computed": true,
        "description": "This unique identifier for Contact",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "is_sensitive_events_subscribed": {
        "computed": true,
        "description": "Whether the account contact association is subscribed to sensitive events. Access to sensitive events is gated by the SubscribeSensitiveEvents virtual IAM action.",
        "description_kind": "plain",
        "type": "bool"
      },
      "managed_notification_configuration_arn": {
        "computed": true,
        "description": "The managed notification configuration ARN, against which the account contact association will be created",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Notifications::ManagedNotificationAccountContactAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNotificationsManagedNotificationAccountContactAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsManagedNotificationAccountContactAssociation), &result)
	return &result
}
