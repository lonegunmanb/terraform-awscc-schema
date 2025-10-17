package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccServicecatalogLaunchNotificationConstraint = `{
  "block": {
    "attributes": {
      "accept_language": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "launch_notification_constraint_id": {
        "computed": true,
        "description": "Unique identifier for the constraint",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_arns": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "portfolio_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "product_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::ServiceCatalog::LaunchNotificationConstraint",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccServicecatalogLaunchNotificationConstraintSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccServicecatalogLaunchNotificationConstraint), &result)
	return &result
}
