package data

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
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identifier": {
        "computed": true,
        "description": "The ID of the Amazon DataZone environment action.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the environment action.",
        "description_kind": "plain",
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
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::DataZone::EnvironmentActions",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneEnvironmentActionsSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneEnvironmentActions), &result)
	return &result
}
