package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftAlias = `{
  "block": {
    "attributes": {
      "alias_id": {
        "computed": true,
        "description": "Unique alias ID",
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description": "A human-readable description of the alias.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "A descriptive label that is associated with an alias. Alias names do not need to be unique.",
        "description_kind": "plain",
        "type": "string"
      },
      "routing_strategy": {
        "computed": true,
        "description": "A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fleet_id": {
              "computed": true,
              "description": "A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.",
              "description_kind": "plain",
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.",
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      }
    },
    "description": "Data Source schema for AWS::GameLift::Alias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccGameliftAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftAlias), &result)
	return &result
}
