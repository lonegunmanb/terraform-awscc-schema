package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamRolePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The policy document.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "The friendly name (not ARN) identifying the policy.",
        "description_kind": "plain",
        "type": "string"
      },
      "role_name": {
        "computed": true,
        "description": "The name of the policy document.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IAM::RolePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamRolePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamRolePolicy), &result)
	return &result
}
