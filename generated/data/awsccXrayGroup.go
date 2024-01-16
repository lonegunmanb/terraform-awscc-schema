package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccXrayGroup = `{
  "block": {
    "attributes": {
      "filter_expression": {
        "computed": true,
        "description": "The filter expression defining criteria by which to group traces.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_arn": {
        "computed": true,
        "description": "The ARN of the group that was generated on creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "group_name": {
        "computed": true,
        "description": "The case-sensitive name of the new group. Names must be unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "insights_configuration": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "insights_enabled": {
              "computed": true,
              "description": "Set the InsightsEnabled value to true to enable insights or false to disable insights.",
              "description_kind": "plain",
              "type": "bool"
            },
            "notifications_enabled": {
              "computed": true,
              "description": "Set the NotificationsEnabled value to true to enable insights notifications. Notifications can only be enabled on a group with InsightsEnabled set to true.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "nesting_mode": "single"
        }
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::XRay::Group",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccXrayGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccXrayGroup), &result)
	return &result
}
