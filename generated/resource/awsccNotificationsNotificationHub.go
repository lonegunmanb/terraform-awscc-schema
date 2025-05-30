package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsNotificationHub = `{
  "block": {
    "attributes": {
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_hub_status_summary": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "notification_hub_status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "notification_hub_status_reason": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "region": {
        "description": "Region that NotificationHub is present in.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::Notifications::NotificationHub",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccNotificationsNotificationHubSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsNotificationHub), &result)
	return &result
}
