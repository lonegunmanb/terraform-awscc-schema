package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccLogsAccountPolicy = `{
  "block": {
    "attributes": {
      "account_id": {
        "computed": true,
        "description": "User account id",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "description": "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per PolicyType.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the account policy",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_type": {
        "description": "Type of the policy.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope for policy application",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description": "The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccLogsAccountPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsAccountPolicy), &result)
	return &result
}
