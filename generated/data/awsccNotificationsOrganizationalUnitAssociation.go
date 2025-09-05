package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsOrganizationalUnitAssociation = `{
  "block": {
    "attributes": {
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
      },
      "organizational_unit_id": {
        "computed": true,
        "description": "The ID of the organizational unit.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Notifications::OrganizationalUnitAssociation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNotificationsOrganizationalUnitAssociationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsOrganizationalUnitAssociation), &result)
	return &result
}
