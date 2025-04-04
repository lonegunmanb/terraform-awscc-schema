package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccGameliftAlias = `{
  "block": {
    "attributes": {
      "alias_arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) that is assigned to a Amazon GameLift Alias resource and uniquely identifies it. ARNs are unique across all Regions. In a GameLift Alias ARN, the resource ID matches the AliasId value.",
        "description_kind": "plain",
        "type": "string"
      },
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
        "description": "A descriptive label that is associated with an alias. Alias names do not need to be unique.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "routing_strategy": {
        "description": "A routing configuration that specifies where traffic is directed for this alias, such as to a fleet or to a message.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "fleet_id": {
              "computed": true,
              "description": "A unique identifier for a fleet that the alias points to. If you specify SIMPLE for the Type property, you must specify this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "message": {
              "computed": true,
              "description": "The message text to be used with a terminal routing strategy. If you specify TERMINAL for the Type property, you must specify this property.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "type": {
              "description": "Simple routing strategy. The alias resolves to one specific fleet. Use this type when routing to active fleets.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "nesting_mode": "single"
        },
        "required": true
      },
      "tags": {
        "computed": true,
        "description": "An array of key-value pairs to apply to this resource.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length.",
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
    "description": "The AWS::GameLift::Alias resource creates an alias for an Amazon GameLift (GameLift) fleet destination.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccGameliftAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccGameliftAlias), &result)
	return &result
}
