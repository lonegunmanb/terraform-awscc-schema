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
