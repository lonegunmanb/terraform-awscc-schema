package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIdentitystoreGroup = `{
  "block": {
    "attributes": {
      "description": {
        "computed": true,
        "description": "A string containing the description of the group.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description": "A string containing the name of the group. This value is commonly displayed when the group is referenced.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_id": {
        "computed": true,
        "description": "The unique identifier for a group in the identity store.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "identity_store_id": {
        "description": "The globally unique identifier for the identity store.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Resource Type definition for AWS::IdentityStore::Group",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIdentitystoreGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIdentitystoreGroup), &result)
	return &result
}
