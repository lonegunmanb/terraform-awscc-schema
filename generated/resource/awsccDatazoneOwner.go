package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccDatazoneOwner = `{
  "block": {
    "attributes": {
      "domain_identifier": {
        "description": "The ID of the domain in which you want to add the entity owner.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_identifier": {
        "description": "The ID of the entity to which you want to add an owner.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "entity_type": {
        "description": "The type of an entity.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "owner": {
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
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
                    "optional": true,
                    "type": "string"
                  }
                },
                "nesting_mode": "single"
              },
              "optional": true
            }
          },
          "nesting_mode": "single"
        },
        "required": true
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
    "description": "A owner can set up authorization permissions on their resources.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccDatazoneOwnerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccDatazoneOwner), &result)
	return &result
}
