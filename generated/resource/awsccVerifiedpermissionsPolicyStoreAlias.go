package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsPolicyStoreAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
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
      "policy_store_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Definition of AWS::VerifiedPermissions::PolicyStoreAlias Resource Type",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccVerifiedpermissionsPolicyStoreAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicyStoreAlias), &result)
	return &result
}
