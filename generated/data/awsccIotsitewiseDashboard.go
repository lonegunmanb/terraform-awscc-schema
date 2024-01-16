package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIotsitewiseDashboard = `{
  "block": {
    "attributes": {
      "dashboard_arn": {
        "computed": true,
        "description": "The ARN of the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_definition": {
        "computed": true,
        "description": "The dashboard definition specified in a JSON literal.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_description": {
        "computed": true,
        "description": "A description for the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_id": {
        "computed": true,
        "description": "The ID of the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_name": {
        "computed": true,
        "description": "A friendly name for the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "project_id": {
        "computed": true,
        "description": "The ID of the project in which to create the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of key-value pairs that contain metadata for the dashboard.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      }
    },
    "description": "Data Source schema for AWS::IoTSiteWise::Dashboard",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIotsitewiseDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIotsitewiseDashboard), &result)
	return &result
}
