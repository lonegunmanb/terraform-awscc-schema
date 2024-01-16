package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSupportappAccountAlias = `{
  "block": {
    "attributes": {
      "account_alias": {
        "computed": true,
        "description": "An account alias associated with a customer's account.",
        "description_kind": "plain",
        "type": "string"
      },
      "account_alias_resource_id": {
        "computed": true,
        "description": "Unique identifier representing an alias tied to an account",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SupportApp::AccountAlias",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSupportappAccountAliasSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSupportappAccountAlias), &result)
	return &result
}
