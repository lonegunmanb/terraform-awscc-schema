package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsManagedNotificationAccountContactAssociation = `{
  "block": {
    "attributes": {
      "contact_identifier": {
        "description": "This unique identifier for Contact",
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
        "description": "The managed notification configuration ARN, against which the account contact association will be created",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "This resource schema represents the ManagedNotificationAccountContactAssociation resource in the AWS User Notifications.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNotificationsManagedNotificationAccountContactAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsManagedNotificationAccountContactAssociation), &result)
	return &result
}
