package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSecurityhubDelegatedAdmin = `{
  "block": {
    "attributes": {
      "admin_account_id": {
        "description": "The AWS-account identifier of the account to designate as the Security Hub CSPM administrator account.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delegated_admin_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "The ` + "`" + `` + "`" + `AWS::SecurityHub::DelegatedAdmin` + "`" + `` + "`" + ` resource designates the delegated ASHlong administrator account for an organization. You must enable the integration between ASH and AOlong before you can designate a delegated ASH administrator. Only the management account for an organization can designate the delegated ASH administrator account. For more information, see [Designating the delegated administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in the *User Guide*.\n To change the delegated administrator account, remove the current delegated administrator account, and then designate the new account.\n To designate multiple delegated administrators in different organizations and AWS-Regions, we recommend using [mappings](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/mappings-section-structure.html).\n Tags aren't supported for this resource.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccSecurityhubDelegatedAdminSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSecurityhubDelegatedAdmin), &result)
	return &result
}
