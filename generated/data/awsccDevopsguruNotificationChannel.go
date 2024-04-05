package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDevopsguruNotificationChannel = `{
  "block": {
    "attributes": {
      "config": {
        "computed": true,
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
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "severities": {
                    "computed": true,
                    "description": "DevOps Guru insight severities to filter for",
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "nesting_mode": "single"
              }
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
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_channel_id": {
        "computed": true,
        "description": "The ID of a notification channel.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DevOpsGuru::NotificationChannel",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDevopsguruNotificationChannelSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDevopsguruNotificationChannel), &result)
	return &result
}
