package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamGroupPolicy = `{
  "block": {
    "attributes": {
      "group_name": {
        "computed": true,
        "description": "The name of the group to associate the policy with.",
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
        "description": "The policy document.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_name": {
        "computed": true,
        "description": "The name of the policy document.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IAM::GroupPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamGroupPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamGroupPolicy), &result)
	return &result
}
