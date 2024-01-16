package data

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
        "type": "string"
      },
      "detail_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_type_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_type_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "target_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "targets": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "target_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::CodeStarNotifications::NotificationRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCodestarnotificationsNotificationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCodestarnotificationsNotificationRule), &result)
	return &result
}
