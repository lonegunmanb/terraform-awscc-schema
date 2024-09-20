package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneEnvironmentActions = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "The description of the Amazon DataZone environment action.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the environment is created.",
        "description_kind": "plain",
        "type": "string"
      },
      "domain_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone domain in which the environment would be created.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "environment_actions_id": {
        "computed": true,
        "description": "The ID of the Amazon DataZone environment action.",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_id": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone environment in which the action is taking place",
        "description_kind": "plain",
        "type": "string"
      },
      "environment_identifier": {
        "computed": true,
        "description": "The identifier of the Amazon DataZone environment in which the action is taking place",
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
      "identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone environment action.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the environment action.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "computed": true,
        "description": "The parameters of the environment action.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "uri": {
              "computed": true,
              "description": "The URI of the console link specified as part of the environment action.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      }
    },
    "description": "Definition of AWS::DataZone::EnvironmentActions Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneEnvironmentActionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironmentActions), &result)
	return &result
}
