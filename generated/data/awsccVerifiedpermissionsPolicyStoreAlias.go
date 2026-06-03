package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccVerifiedpermissionsPolicyStoreAlias = `{
  "block": {
    "attributes": {
      "alias_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_store_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::VerifiedPermissions::PolicyStoreAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccVerifiedpermissionsPolicyStoreAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccVerifiedpermissionsPolicyStoreAlias), &result)
	return &result
}
