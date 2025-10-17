package resource

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
        "description": "The name of the Automation document to execute",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "automation_parameters": {
        "description": "Parameters to be passed to the Automation Document",
        "description_kind": "plain",
        "required": true,
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_key": {
        "description": "A unique identifier used for generating a unique logical ID for the custom resource",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "description": "Resource Type definition for AWS::SSMQuickSetup::LifecycleAutomation that executes SSM Automation documents in response to CloudFormation lifecycle events.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSsmquicksetupLifecycleAutomationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmquicksetupLifecycleAutomation), &result)
	return &result
}
