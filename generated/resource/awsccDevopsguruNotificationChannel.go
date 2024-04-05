package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsguruNotificationChannel = `{
  "block": {
    "attributes": {
      "config": {
        "description": "Information about notification channels you have configured with DevOps Guru.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "filters": {
              "computed": true,
              "description": "Information about filters of a notification channel configured in DevOpsGuru to filter for insights.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "message_types": {
                    "computed": true,
                    "description": "DevOps Guru message types to filter for",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "severities": {
                    "computed": true,
                    "description": "DevOps Guru insight severities to filter for",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "sns": {
              "computed": true,
              "description": "Information about a notification channel configured in DevOps Guru to send notifications when insights are created.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "topic_arn": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_channel_id": {
        "computed": true,
        "description": "The ID of a notification channel.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "This resource schema represents the NotificationChannel resource in the Amazon DevOps Guru.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDevopsguruNotificationChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruNotificationChannel), &result)
	return &result
}
