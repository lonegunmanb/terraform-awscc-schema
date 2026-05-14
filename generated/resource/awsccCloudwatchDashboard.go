package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchDashboard = `{
  "block": {
    "attributes": {
      "dashboard_body": {
        "description": "The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "dashboard_name": {
        "computed": true,
        "description": "The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs to associate with the cloudwatch dashboard. You can associate up to 50 tags with a dashboard",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "A unique identifier for the tag. The combination of tag keys and values can help you organize and categorize your resources.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the specified tag key.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      }
    },
    "description": "Resource Type definition for AWS::CloudWatch::Dashboard",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDashboard), &result)
	return &result
}
