package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportappAccountAlias = `{
  "block": {
    "attributes": {
      "account_alias": {
        "description": "An account alias associated with a customer's account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_alias_resource_id": {
        "computed": true,
        "description": "Unique identifier representing an alias tied to an account",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "An AWS Support App resource that creates, updates, reads, and deletes a customer's account alias.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSupportappAccountAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportappAccountAlias), &result)
	return &result
}
