package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccOrganizationsOrganizationalUnit = `{
  "block": {
    "attributes": {
      "arn": {
        "computed": true,
        "description": "The Amazon Resource Name (ARN) of this OU.",
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
        "description": "The friendly name of this OU.",
        "description_kind": "plain",
        "type": "string"
      },
      "organizational_unit_id": {
        "computed": true,
        "description": "The unique identifier (ID) associated with this OU.",
        "description_kind": "plain",
        "type": "string"
      },
      "parent_id": {
        "computed": true,
        "description": "The unique identifier (ID) of the parent root or OU that you want to create the new OU in.",
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description": "A list of tags that you want to attach to the newly created OU.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "key": {
              "computed": true,
              "description": "The key identifier, or name, of the tag.",
              "description_kind": "plain",
              "type": "string"
            },
            "value": {
              "computed": true,
              "description": "The string value that's associated with the key of the tag. You can set the value of a tag to an empty string, but you can't set the value of a tag to null.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "set"
        }
      }
    },
    "description": "Data Source schema for AWS::Organizations::OrganizationalUnit",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccOrganizationsOrganizationalUnitSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccOrganizationsOrganizationalUnit), &result)
	return &result
}
