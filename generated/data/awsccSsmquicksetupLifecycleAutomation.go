package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmquicksetupLifecycleAutomation = `{
  "block": {
    "attributes": {
      "association_id": {
        "computed": true,
        "description": "The id from the association that is returned when creating the association",
        "description_kind": "plain",
        "type": "string"
      },
      "automation_document": {
        "computed": true,
        "description": "The name of the Automation document to execute",
        "description_kind": "plain",
        "type": "string"
      },
      "automation_parameters": {
        "computed": true,
        "description": "Parameters to be passed to the Automation Document",
        "description_kind": "plain",
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_key": {
        "computed": true,
        "description": "A unique identifier used for generating a unique logical ID for the custom resource",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Data Source schema for AWS::SSMQuickSetup::LifecycleAutomation",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmquicksetupLifecycleAutomationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmquicksetupLifecycleAutomation), &result)
	return &result
}
