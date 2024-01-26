package data

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
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The body of the policy document you want to use for this topic.\n\nYou can only add one policy per PolicyType.\n\nThe policy must be in JSON string format.\n\nLength Constraints: Maximum length of 30720",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "The name of the account policy",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_type": {
        "computed": true,
        "description": "Type of the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "scope": {
        "computed": true,
        "description": "Scope for policy application",
        "description_kind": "plain",
        "type": "string"
      },
      "selection_criteria": {
        "computed": true,
        "description": "Log group  selection criteria to apply policy only to a subset of log groups. SelectionCriteria string can be up to 25KB and cloudwatchlogs determines the length of selectionCriteria by using its UTF-8 bytes",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::Logs::AccountPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccLogsAccountPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccLogsAccountPolicy), &result)
	return &result
}
