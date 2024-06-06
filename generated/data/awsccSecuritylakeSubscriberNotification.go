package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecuritylakeSubscriberNotification = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "https_notification_configuration": {
              "computed": true,
              "description": "The configuration for HTTPS subscriber notification.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "authorization_api_key_name": {
                    "computed": true,
                    "description": "The key name for the notification subscription.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "authorization_api_key_value": {
                    "computed": true,
                    "description": "The key value for the notification subscription.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "endpoint": {
                    "computed": true,
                    "description": "The subscription endpoint in Security Lake.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "http_method": {
                    "computed": true,
                    "description": "The HTTPS method used for the notification subscription.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "target_role_arn": {
                    "computed": true,
                    "description": "The Amazon Resource Name (ARN) of the EventBridge API destinations IAM role that you created.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "sqs_notification_configuration": {
              "computed": true,
              "description": "The configurations for SQS subscriber notification. The members of this structure are context-dependent.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "subscriber_arn": {
        "computed": true,
        "description": "The ARN for the subscriber",
        "description_kind": "plain",
        "type": "string"
      },
      "subscriber_endpoint": {
        "computed": true,
        "description": "The endpoint the subscriber should listen to for notifications",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityLake::SubscriberNotification",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecuritylakeSubscriberNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecuritylakeSubscriberNotification), &result)
	return &result
}
