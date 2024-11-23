package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudtrailDashboard = `{
  "block": {
    "attributes": {
      "created_timestamp": {
        "computed": true,
        "description": "The timestamp of the dashboard creation.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_arn": {
        "computed": true,
        "description": "The ARN of the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the dashboard.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "refresh_schedule": {
        "computed": true,
        "description": "Configures the automatic refresh schedule for the dashboard. Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "frequency": {
              "computed": true,
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "unit": {
                    "computed": true,
                    "description": "The frequency unit. Supported values are HOURS and DAYS.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "value": {
                    "computed": true,
                    "description": "The frequency value.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            },
            "status": {
              "computed": true,
              "description": "The status of the schedule. Supported values are ENABLED and DISABLED.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "time_of_day": {
              "computed": true,
              "description": "StartTime of the automatic schedule refresh.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "status": {
        "computed": true,
        "description": "The status of the dashboard. Values are CREATING, CREATED, UPDATING, UPDATED and DELETING.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      },
      "termination_protection_enabled": {
        "computed": true,
        "description": "Indicates whether the dashboard is protected from termination.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "type": {
        "computed": true,
        "description": "The type of the dashboard. Values are CUSTOM and MANAGED.",
        "description_kind": "plain",
        "type": "string"
      },
      "updated_timestamp": {
        "computed": true,
        "description": "The timestamp showing when the dashboard was updated, if applicable. UpdatedTimestamp is always either the same or newer than the time shown in CreatedTimestamp.",
        "description_kind": "plain",
        "type": "string"
      },
      "widgets": {
        "computed": true,
        "description": "List of widgets on the dashboard",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "query_parameters": {
              "computed": true,
              "description": "The placeholder keys in the QueryStatement. For example: $StartTime$, $EndTime$, $Period$.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_statement": {
              "computed": true,
              "description": "The SQL query statement on one or more event data stores.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "view_properties": {
              "computed": true,
              "description": "The view properties of the widget.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "nesting_mode": "list"
        },
        "optional": true
      }
    },
    "description": "The Amazon CloudTrail dashboard resource allows customers to manage managed dashboards and create custom dashboards. You can manually refresh custom and managed dashboards. For custom dashboards, you can also set up an automatic refresh schedule and modify dashboard widgets.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudtrailDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudtrailDashboard), &result)
	return &result
}
