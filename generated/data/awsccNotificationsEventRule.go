package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccNotificationsEventRule = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "creation_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_pattern": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_type": {
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
      "managed_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "notification_configuration_arn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "regions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "source": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status_summary_by_region": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "reason": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "map"
        }
      }
    },
    "description": "Data Source schema for AWS::Notifications::EventRule",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccNotificationsEventRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccNotificationsEventRule), &result)
	return &result
}
