package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccConnectNotification = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) for the notification.",
        "description_kind": "plain",
        "type": "string"
      },
      "content": {
        "computed": true,
        "description": "The content of the notification.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "de_de": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "en_us": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "es_es": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "fr_fr": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "id_id": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "it_it": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "ja_jp": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "ko_kr": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "pt_br": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "zh_cn": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            },
            "zh_tw": {
              "computed": true,
              "description": "Localized notification content",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "created_at": {
        "computed": true,
        "description": "The time a notification was created",
        "description_kind": "plain",
        "type": "string"
      },
      "expires_at": {
        "computed": true,
        "description": "The time a notification will expire",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "instance_arn": {
        "computed": true,
        "description": "The identifier of the Amazon Connect instance.",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_id": {
        "computed": true,
        "description": "The identifier of the notification.",
        "description_kind": "plain",
        "type": "string"
      },
      "priority": {
        "computed": true,
        "description": "The priority of the notification.",
        "description_kind": "plain",
        "type": "string"
      },
      "recipients": {
        "computed": true,
        "description": "The recipients of the notification.",
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "computed": true,
        "description": "One or more tags.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Connect::Notification",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccConnectNotificationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccConnectNotification), &result)
	return &result
}
