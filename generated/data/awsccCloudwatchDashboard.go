package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchDashboard = `{
  "block": {
    "attributes": {
      "dashboard_body": {
        "computed": true,
        "description": "The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_name": {
        "computed": true,
        "description": "The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::CloudWatch::Dashboard",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccCloudwatchDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchDashboard), &result)
	return &result
}
