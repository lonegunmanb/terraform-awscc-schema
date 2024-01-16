package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccSsmResourcePolicy = `{
  "block": {
    "attributes": {
      "id": {
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description": "Actual policy statement.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_hash": {
        "computed": true,
        "description": "A snapshot identifier for the policy over time.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_id": {
        "computed": true,
        "description": "An unique identifier within the policies of a resource. ",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_arn": {
        "computed": true,
        "description": "Arn of OpsItemGroup etc.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::SSM::ResourcePolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccSsmResourcePolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccSsmResourcePolicy), &result)
	return &result
}
