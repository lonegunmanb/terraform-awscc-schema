package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIdentitystoreGroupMembership = `{
  "block": {
    "attributes": {
      "group_id": {
        "computed": true,
        "description": "The unique identifier for a group in the identity store.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "identity_store_id": {
        "computed": true,
        "description": "The globally unique identifier for the identity store.",
        "description_kind": "plain",
        "type": "string"
      },
      "member_id": {
        "computed": true,
        "description": "An object containing the identifier of a group member.",
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "user_id": {
              "computed": true,
              "description": "The identifier for a user in the identity store.",
              "description_kind": "plain",
              "type": "string"
            }
          },
          "nesting_mode": "single"
        }
      },
      "membership_id": {
        "computed": true,
        "description": "The identifier for a GroupMembership in the identity store.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IdentityStore::GroupMembership",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIdentitystoreGroupMembershipSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIdentitystoreGroupMembership), &result)
	return &result
}
