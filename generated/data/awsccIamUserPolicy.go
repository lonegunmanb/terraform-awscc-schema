package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamUserPolicy = `{
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
        "description": "The name of the policy document.",
        "description_kind": "plain",
        "type": "string"
      },
      "user_name": {
        "computed": true,
        "description": "The name of the user to associate the policy with.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "description": "Data Source schema for AWS::IAM::UserPolicy",
    "description_kind": "plain"
  },
  "version": 0
}`

func AwsccIamUserPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamUserPolicy), &result)
	return &result
}
