package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneOwner = `{
  "block": {
    "attributes": {
      "domain_identifier": {
        "computed": true,
        "description": "The ID of the domain in which you want to add the entity owner.",
        "description_kind": "plain",
        "type": "string"
      },
      "entity_identifier": {
        "computed": true,
        "description": "The ID of the entity to which you want to add an owner.",
        "description_kind": "plain",
        "type": "string"
      },
      "entity_type": {
        "computed": true,
        "description": "The type of an entity.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "owner": {
        "computed": true,
        "description": "The owner that you want to add to the entity.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "group": {
              "computed": true,
              "description": "The properties of the domain unit owners group.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "group_identifier": {
                    "computed": true,
                    "description": "The ID of the domain unit owners group.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            },
            "user": {
              "computed": true,
              "description": "The properties of the owner user.",
              "description_kind": "plain",
              "nested_type": {
                "attributes": {
                  "user_identifier": {
                    "computed": true,
                    "description": "The ID of the owner user.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              }
            }
          },
          "nesting_mode": "single"
        }
      },
      "owner_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "owner_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::DataZone::Owner",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccDatazoneOwnerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneOwner), &result)
	return &result
}
