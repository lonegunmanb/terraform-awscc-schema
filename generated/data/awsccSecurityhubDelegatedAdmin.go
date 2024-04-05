package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubDelegatedAdmin = `{
  "block": {
    "attributes": {
      "admin_account_id": {
        "computed": true,
        "description": "The Amazon Web Services account identifier of the account to designate as the Security Hub administrator account",
        "description_kind": "plain",
        "type": "string"
      },
      "delegated_admin_identifier": {
        "computed": true,
        "description": "The identifier of the DelegatedAdmin being created and assigned as the unique identifier",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "status": {
        "computed": true,
        "description": "The current status of the Security Hub administrator account. Indicates whether the account is currently enabled as a Security Hub administrator",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SecurityHub::DelegatedAdmin",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSecurityhubDelegatedAdminSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubDelegatedAdmin), &result)
	return &result
}
