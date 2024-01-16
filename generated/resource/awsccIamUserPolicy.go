package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamUserPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "Uniquely identifies the resource.",
        "description_kind": "plain",
        "type": "string"
      },
      "policy_document": {
        "computed": true,
        "description": "The policy document.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_name": {
        "description": "The name of the policy document.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "user_name": {
        "description": "The name of the user to associate the policy with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "description": "Schema for IAM User Policy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamUserPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamUserPolicy), &result)
	return &result
}
