package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const awsccIamGroupPolicy = `{
  "block": {
    "attributes": {
      "group_name": {
        "description": "The name of the group to associate the policy with.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
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
      }
    },
    "description": "Schema for IAM Group Policy",
    "description_kind": "plain"
  },
  "version": 1
}`

func AwsccIamGroupPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(awsccIamGroupPolicy), &result)
	return &result
}
