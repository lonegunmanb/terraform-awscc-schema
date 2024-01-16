package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCodestarnotificationsNotificationRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "detail_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_type_id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "event_type_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "target_address": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "targets": {
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "required": true
      }
    },
    "description": "Resource Type definition for AWS::CodeStarNotifications::NotificationRule",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCodestarnotificationsNotificationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodestarnotificationsNotificationRule), &result)
	return &result
}
