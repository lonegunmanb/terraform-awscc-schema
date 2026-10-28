package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccCloudwatchOmniDashboard = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "body": {
        "description": "The dashboard body, a JSON document describing the dashboard's widgets and layout.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "created_at": {
        "computed": true,
        "description": "The time the dashboard was created.",
        "description_kind": "plain",
        "type": "string"
      },
      "created_by": {
        "computed": true,
        "description": "The principal that created the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "dashboard_id": {
        "computed": true,
        "description": "The unique identifier of the dashboard.",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "The description of the dashboard.",
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
      "name": {
        "description": "The name of the dashboard.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "space_id": {
        "description": "The unique identifier of the space the dashboard belongs to.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "set"
        },
        "optional": true
      },
      "updated_at": {
        "computed": true,
        "description": "The time the dashboard was last updated.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Resource type definition for AWS::CloudWatch::OmniDashboard",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccCloudwatchOmniDashboardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccCloudwatchOmniDashboard), &result)
	return &result
}
